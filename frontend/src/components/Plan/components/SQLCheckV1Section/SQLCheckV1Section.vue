<template>
  <div v-if="show" class="px-4 pt-3 flex flex-col gap-y-1 overflow-hidden">
    <div class="flex items-center justify-between gap-2">
      <div class="flex items-center gap-2">
        <h3 class="text-base font-medium">{{ $t("plan.checks.self") }}</h3>

        <NTooltip v-if="checkResults && affectedRows > 0">
          <template #trigger>
            <NTag round :bordered="false">
              <span class="text-sm text-control-light mr-1">{{
                $t("task.check-type.affected-rows.self")
              }}</span>
              <span class="text-sm font-medium">
                {{ affectedRows }}
              </span>
            </NTag>
          </template>
          {{ $t("task.check-type.affected-rows.description") }}
        </NTooltip>
      </div>

      <div class="flex items-center gap-4">
        <!-- Status Summary -->
        <div class="flex items-center gap-2 text-sm">
          <button
            v-if="summary.error > 0"
            class="flex items-center gap-1 hover:opacity-80"
            @click="openDrawer('ERROR')"
          >
            <XCircleIcon class="w-5 h-5 text-error" />
            <span class="text-error font-medium">{{ summary.error }}</span>
          </button>
          <button
            v-if="summary.warning > 0"
            class="flex items-center gap-1 hover:opacity-80"
            @click="openDrawer('WARNING')"
          >
            <AlertCircleIcon class="w-5 h-5 text-warning" />
            <span class="text-warning font-medium">{{ summary.warning }}</span>
          </button>
          <span
            v-if="
              !isNullOrUndefined(checkResults) &&
              summary.error === 0 &&
              summary.warning === 0
            "
            class="flex items-center"
          >
            <CheckCircleIcon class="w-5 h-5 text-success" />
          </span>
        </div>

        <!-- Run Checks Button -->
        <NButton
          size="small"
          :loading="isRunningChecks"
          :disabled="statement.length === 0"
          @click="runChecks"
        >
          <template #icon>
            <PlayIcon class="w-4 h-4" />
          </template>
          {{ $t("task.run-checks") }}
        </NButton>
      </div>
    </div>

    <!-- Status Drawer -->
    <Drawer v-model:show="drawerVisible">
      <DrawerContent
        :title="$t('plan.navigator.checks')"
        class="w-[40rem] max-w-[100vw] relative"
      >
        <div class="w-full h-full flex flex-col">
          <!-- Drawer Content -->
          <div v-if="drawerAdvices.length > 0" class="w-full space-y-2">
            <div
              v-for="(advice, idx) in drawerAdvices"
              :key="idx"
              class="space-y-1 p-3 border rounded-lg bg-gray-50"
            >
              <div class="flex items-start justify-between">
                <div class="text-sm font-medium text-main">
                  {{ getAdviceTitle(advice) }}
                </div>
                <component
                  :is="getStatusIcon(advice.status)"
                  class="w-4 h-4 flex-shrink-0"
                  :class="getStatusColor(advice.status)"
                />
              </div>

              <!-- Target Database -->
              <DatabaseDisplay :database="advice.target" />

              <!-- Advice Content -->
              <div v-if="advice.content" class="text-xs text-control-light">
                {{ advice.content }}
              </div>

              <!-- Location Info -->
              <div
                v-if="advice.startPosition && advice.startPosition.line > 0"
                class="text-xs text-control-lighter"
              >
                Line {{ advice.startPosition.line }}, Column
                {{ advice.startPosition.column }}
              </div>
            </div>
          </div>
          <div v-else class="w-full text-center py-8 text-control-light">
            No check results
          </div>
        </div>
      </DrawerContent>
    </Drawer>
  </div>
</template>

<script setup lang="ts">
import { create } from "@bufbuild/protobuf";
import {
  CheckCircleIcon,
  AlertCircleIcon,
  XCircleIcon,
  PlayIcon,
} from "lucide-vue-next";
import { NButton, NTag, NTooltip } from "naive-ui";
import { computed, ref, watch } from "vue";
import { getLocalSheetByName } from "@/components/Plan";
import Drawer from "@/components/v2/Container/Drawer.vue";
import DrawerContent from "@/components/v2/Container/DrawerContent.vue";
import { releaseServiceClientConnect } from "@/grpcweb";
import { projectNamePrefix } from "@/store";
import { getRuleLocalization, ruleTemplateMapV2 } from "@/types";
import { Plan_ChangeDatabaseConfig_Type } from "@/types/proto-es/v1/plan_service_pb";
import {
  CheckReleaseRequestSchema,
  Release_File_Type,
} from "@/types/proto-es/v1/release_service_pb";
import {
  Release_File_ChangeType,
  type CheckReleaseResponse_CheckResult,
} from "@/types/proto-es/v1/release_service_pb";
import { Advice_Status, type Advice } from "@/types/proto-es/v1/sql_service_pb";
import {
  extractProjectResourceName,
  getSheetStatement,
  isNullOrUndefined,
} from "@/utils";
import { usePlanContext } from "../../logic/context";
import { targetsForSpec } from "../../logic/plan";
import { useSelectedSpec } from "../SpecDetailView/context";
import DatabaseDisplay from "../common/DatabaseDisplay.vue";

const { plan } = usePlanContext();
const selectedSpec = useSelectedSpec();

const isRunningChecks = ref(false);
const drawerVisible = ref(false);
const selectedStatus = ref<"ERROR" | "WARNING" | "SUCCESS">("ERROR");
const checkResults = ref<CheckReleaseResponse_CheckResult[] | undefined>(
  undefined
);

const statement = computed(() => {
  if (!selectedSpec.value) return "";
  const config =
    selectedSpec.value.config?.case === "changeDatabaseConfig"
      ? selectedSpec.value.config.value
      : undefined;
  if (!config) return "";
  const sheet = getLocalSheetByName(config.sheet);
  return getSheetStatement(sheet);
});

const show = computed(() => {
  if (!selectedSpec.value) {
    return false;
  }
  // Show for change database configs
  return selectedSpec.value.config?.case === "changeDatabaseConfig";
});

// Enhanced advice type with target information
type AdviceWithTarget = Advice & {
  target: string;
};

const allAdvices = computed(() => {
  if (!checkResults.value) {
    return [];
  }
  const advices: AdviceWithTarget[] = [];
  for (const result of checkResults.value) {
    for (const advice of result.advices) {
      advices.push({
        ...advice,
        target: result.target,
      });
    }
  }
  return advices;
});

const summary = computed(() => {
  const result = { success: 0, warning: 0, error: 0 };

  for (const advice of allAdvices.value) {
    if (advice.status === Advice_Status.ERROR) {
      result.error++;
    } else if (advice.status === Advice_Status.WARNING) {
      result.warning++;
    } else if (advice.status === Advice_Status.SUCCESS) {
      result.success++;
    }
  }

  return result;
});

const drawerAdvices = computed(() => {
  return allAdvices.value.filter((advice) => {
    if (selectedStatus.value === "ERROR") {
      return advice.status === Advice_Status.ERROR;
    } else if (selectedStatus.value === "WARNING") {
      return advice.status === Advice_Status.WARNING;
    } else {
      return advice.status === Advice_Status.SUCCESS;
    }
  });
});

const runChecks = async () => {
  if (!plan.value.name || !selectedSpec.value) return;

  const config =
    selectedSpec.value.config?.case === "changeDatabaseConfig"
      ? selectedSpec.value.config.value
      : undefined;
  if (!config) return;

  isRunningChecks.value = true;
  try {
    // Get the statement from the sheet
    const sheetName = config.sheet;
    const sheet = getLocalSheetByName(sheetName);
    const statement = getSheetStatement(sheet);

    // Get targets
    const targets = targetsForSpec(selectedSpec.value);

    // Run check for all targets
    const request = create(CheckReleaseRequestSchema, {
      parent: `${projectNamePrefix}${extractProjectResourceName(plan.value.name)}`,
      release: {
        files: [
          {
            // Use "0" for dummy version.
            version: "0",
            type: Release_File_Type.VERSIONED,
            statement: new TextEncoder().encode(statement),
            changeType:
              config.type === Plan_ChangeDatabaseConfig_Type.DATA
                ? Release_File_ChangeType.DML
                : Release_File_ChangeType.DDL,
          },
        ],
      },
      targets,
    });
    const response = await releaseServiceClientConnect.checkRelease(request);
    checkResults.value = response.results || [];
  } finally {
    isRunningChecks.value = false;
  }
};

const affectedRows = computed(() => {
  if (!checkResults.value) return 0;
  return checkResults.value.reduce((acc, result) => {
    return acc + result.affectedRows;
  }, 0);
});

const openDrawer = (status: "ERROR" | "WARNING" | "SUCCESS") => {
  selectedStatus.value = status;
  drawerVisible.value = true;
};

const getStatusIcon = (status: Advice_Status | string) => {
  if (status === Advice_Status.ERROR || status === "ERROR") return XCircleIcon;
  if (status === Advice_Status.WARNING || status === "WARNING")
    return AlertCircleIcon;
  return CheckCircleIcon;
};

const getStatusColor = (status: Advice_Status | string) => {
  if (status === Advice_Status.ERROR || status === "ERROR") return "text-error";
  if (status === Advice_Status.WARNING || status === "WARNING")
    return "text-warning";
  return "text-success";
};

watch(
  () => selectedSpec.value.id,
  () => {
    // Reset drawer and results when spec changes.
    drawerVisible.value = false;
    checkResults.value = undefined;
  },
  {
    immediate: true,
  }
);

const getAdviceTitle = (advice: Advice): string => {
  let title = advice.title;
  const rule = getRuleTemplateByType(advice.title);
  if (rule) {
    const ruleLocalization = getRuleLocalization(rule.type, rule.engine);
    if (ruleLocalization.title) {
      title = ruleLocalization.title;
    }
  }
  return title;
};

const getRuleTemplateByType = (type: string) => {
  for (const mapByType of ruleTemplateMapV2.values()) {
    if (mapByType.has(type)) {
      return mapByType.get(type);
    }
  }
  return;
};
</script>
