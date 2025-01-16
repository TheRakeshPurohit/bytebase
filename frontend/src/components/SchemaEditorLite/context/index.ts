import Emittery from "emittery";
import type { Ref } from "vue";
import { inject, provide, computed } from "vue";
import { supportSetClassificationFromComment } from "@/components/ColumnDataTable/utils";
import { useSettingV1Store } from "@/store";
import type { ComposedProject } from "@/types";
import { Engine } from "@/types/proto/v1/common";
import type { RebuildMetadataEditReset } from "../algorithm/rebuild";
import type { EditTarget, RolloutObject } from "../types";
import { useEditConfigs } from "./config";
import { useEditStatus } from "./edit";
import { useScrollStatus } from "./scroll";
import { useSelection } from "./selection";
import { useTabs } from "./tabs";

export const KEY = Symbol("bb.schema-editor");

export type SchemaEditorEvents = Emittery<{
  ["update:selected-rollout-objects"]: RolloutObject[];
  ["rebuild-tree"]: {
    openFirstChild: boolean;
  };
  ["rebuild-edit-status"]: { resets: RebuildMetadataEditReset[] };
  ["clear-tabs"]: undefined;
}>;

export type SchemaEditorOptions = {
  forceShowIndexes: boolean;
  forceShowPartitions: boolean;
};

export const provideSchemaEditorContext = (params: {
  readonly: Ref<boolean>;
  project: Ref<ComposedProject>;
  targets: Ref<EditTarget[]>;
  selectedRolloutObjects: Ref<RolloutObject[] | undefined>;
  disableDiffColoring: Ref<boolean>;
  hidePreview: Ref<boolean>;
  options?: Ref<SchemaEditorOptions>;
}) => {
  const events = new Emittery() as SchemaEditorEvents;
  const classificationConfig = computed(() => {
    if (!params.project.value.dataClassificationConfigId) {
      return;
    }
    return useSettingV1Store().getProjectClassification(
      params.project.value.dataClassificationConfigId
    );
  });

  const context = {
    events,
    ...params,
    ...useTabs(events),
    ...useEditStatus(),
    ...useEditConfigs(params.targets),
    ...useScrollStatus(),
    ...useSelection(params.selectedRolloutObjects, events),
    classificationConfig,
    showClassificationColumn: (
      engine: Engine,
      classificationFromConfig: boolean
    ) => {
      if (!classificationConfig.value) {
        return false;
      }
      return supportSetClassificationFromComment(
        engine,
        classificationFromConfig
      );
    },
  };

  provide(KEY, context);

  return context;
};

export type SchemaEditorContext = ReturnType<typeof provideSchemaEditorContext>;

export const useSchemaEditorContext = () => {
  return inject(KEY) as SchemaEditorContext;
};
