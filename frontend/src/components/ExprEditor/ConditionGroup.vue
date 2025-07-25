<template>
  <div
    class="w-full overflow-hidden space-y-2 py-0.5"
    :class="[root ? '' : 'border rounded-[3px] bg-gray-50']"
  >
    <div v-if="!root" class="pl-2.5 pr-1 text-gray-500 flex items-center">
      <div class="flex-1">
        <template v-if="args.length > 0">
          <template v-if="operator === '_||_'">
            {{ $t("cel.condition.group.or.description") }}
          </template>
          <template v-if="operator === '_&&_'">
            {{ $t("cel.condition.group.and.description") }}
          </template>
        </template>
        <template v-else>
          <i18n-t
            keypath="cel.condition.add-condition-in-group-placeholder"
            tag="div"
            class="inline-flex items-center"
          >
            <template #plus>
              <heroicons:plus class="w-3 h-3 mx-1" />
            </template>
          </i18n-t>
        </template>
      </div>
      <div v-if="allowAdmin" class="flex items-center justify-end">
        <NButton
          size="tiny"
          quaternary
          type="default"
          :style="`flex-shrink: 0; padding-left: 0; padding-right: 0; --n-width: 22px;`"
          @click="emit('remove')"
        >
          <heroicons:trash class="w-3.5 h-3.5" />
        </NButton>
      </div>
    </div>

    <div v-if="root && args.length === 0" class="px-1.5 text-gray-500">
      {{ $t("cel.condition.add-root-condition-placeholder") }}
    </div>
    <div
      v-for="(operand, i) in args"
      :key="i"
      class="flex items-start gap-x-1 w-full"
      :class="[root ? '' : 'px-1']"
    >
      <div class="w-14 shrink-0">
        <div v-if="i === 0" class="pl-1.5 pt-1 text-control">Where</div>
        <NSelect
          v-else-if="i === 1 && allowAdmin"
          v-model:value="operator"
          :options="OPERATORS"
          :consistent-menu-width="false"
          size="small"
          @update:value="$emit('update')"
        />
        <div v-else class="pl-2 pt-1 text-control lowercase">
          {{ operatorLabel(operator) }}
        </div>
      </div>
      <div class="flex-1 space-y-1 overflow-x-hidden">
        <ConditionGroup
          v-if="isConditionGroupExpr(operand)"
          :expr="operand"
          @remove="removeConditionGroup(operand)"
          @update="$emit('update')"
        />
        <Condition
          v-if="isConditionExpr(operand)"
          :expr="operand"
          @remove="removeCondition(operand)"
          @update="$emit('update')"
        />
        <RawString
          v-if="isRawStringExpr(operand)"
          :expr="operand"
          @remove="removeRawString(operand)"
          @update="$emit('update')"
        />
      </div>
    </div>

    <div v-if="!root && allowAdmin" class="pl-1.5 pb-1">
      <NButton size="small" quaternary @click="addCondition">
        <template #icon
          ><heroicons:plus class="w-4 h-4 text-gray-500"
        /></template>
        <span class="text-gray-500">{{ $t("cel.condition.add") }}</span>
      </NButton>
      <NButton size="small" quaternary @click="addRawString">
        <template #icon
          ><heroicons:plus class="w-4 h-4 text-gray-500"
        /></template>
        <span class="text-gray-500">{{
          $t("cel.condition.add-raw-expression")
        }}</span>
      </NButton>
    </div>

    <div v-if="root && allowAdmin" class="space-x-1">
      <NButton size="small" quaternary @click="addCondition">
        <template #icon><heroicons:plus class="w-4 h-4" /></template>
        <span>{{ $t("cel.condition.add") }}</span>
      </NButton>
      <NButton size="small" quaternary @click="addConditionGroup">
        <template #icon><heroicons:plus class="w-4 h-4" /></template>
        <span>{{ $t("cel.condition.add-group") }}</span>
        <NTooltip>
          <template #trigger>
            <heroicons:question-mark-circle class="ml-1 w-3 h-3" />
          </template>
          <div class="max-w-[18rem]">
            {{ $t("cel.condition.group.tooltip") }}
          </div>
        </NTooltip>
      </NButton>
      <NButton
        v-if="enableRawExpression"
        size="small"
        quaternary
        @click="addRawString"
      >
        <template #icon><heroicons:plus class="w-4 h-4" /></template>
        <span>{{ $t("cel.condition.add-raw-expression") }}</span>
      </NButton>
    </div>
  </div>
</template>

<script lang="ts" setup>
/* eslint-disable vue/no-mutating-props */
import { NButton, NSelect, NTooltip, type SelectOption } from "naive-ui";
import { computed } from "vue";
import {
  type ConditionExpr,
  type ConditionGroupExpr,
  type LogicalOperator,
  LogicalOperatorList,
  isConditionGroupExpr,
  isConditionExpr,
  isRawStringExpr,
  ExprType,
  type RawStringExpr,
  isNumberFactor,
  type Factor,
  isStringFactor,
  isTimestampFactor,
} from "@/plugins/cel";
import Condition from "./Condition.vue";
import RawString from "./RawString.vue";
import { getOperatorListByFactor } from "./components/common";
import { useExprEditorContext } from "./context";

const props = defineProps<{
  expr: ConditionGroupExpr;
  root?: boolean;
}>();

const emit = defineEmits<{
  (event: "remove"): void;
  (event: "update"): void;
}>();

const {
  allowAdmin,
  enableRawExpression,
  factorList,
  factorOperatorOverrideMap,
} = useExprEditorContext();

const operator = computed({
  get() {
    return props.expr.operator;
  },
  set(op) {
    props.expr.operator = op;
  },
});
const args = computed(() => props.expr.args);

const operatorLabel = (op: LogicalOperator) => {
  if (op === "_&&_") return "and";
  if (op === "_||_") return "or";
  throw new Error(`unknown logical operator "${op}"`);
};

const OPERATORS: SelectOption[] = [
  { label: operatorLabel("_&&_"), value: "_&&_" },
  { label: operatorLabel("_||_"), value: "_||_" },
];

const addCondition = () => {
  const factor = factorList.value[0];
  const operators = getOperatorListByFactor(
    factor,
    factorOperatorOverrideMap.value
  );

  args.value.push({
    type: ExprType.Condition,
    operator: operators[0] ?? "",
    args: [factor, getDefaultValue(factor)],
  } as any);
  emit("update");
};

const addRawString = () => {
  args.value.push({
    type: ExprType.RawString,
    content: "",
  });
  emit("update");
};

const addConditionGroup = () => {
  args.value.push({
    type: ExprType.ConditionGroup,
    operator: LogicalOperatorList[0],
    args: [],
  });
  emit("update");
};

const removeCondition = (condition: ConditionExpr) => {
  const index = args.value.indexOf(condition);
  if (index >= 0) {
    args.value.splice(index, 1);
    emit("update");
  }
};

const removeConditionGroup = (group: ConditionGroupExpr) => {
  const index = args.value.indexOf(group);
  if (index >= 0) {
    args.value.splice(index, 1);
    emit("update");
  }
};

const removeRawString = (rawString: RawStringExpr) => {
  const index = args.value.indexOf(rawString);
  if (index >= 0) {
    args.value.splice(index, 1);
    emit("update");
  }
};

const getDefaultValue = (factor: Factor): any => {
  if (isNumberFactor(factor)) return 0;
  if (isStringFactor(factor)) return "";
  if (isTimestampFactor(factor)) return new Date();
  // Fall back to empty string.
  return "";
};
</script>
