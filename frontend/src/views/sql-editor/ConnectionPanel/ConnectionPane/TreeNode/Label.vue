<template>
  <template v-if="type === 'instance'">
    <InstanceNode :node="node" :keyword="keyword" />
  </template>
  <template v-if="type === 'environment'">
    <EnvironmentV1Name
      :environment="(node as TreeNode<'environment'>).meta.target"
      :keyword="keyword"
      :link="false"
    />
  </template>
  <template v-if="type === 'database'">
    <DatabaseNode
      v-bind="$attrs"
      :node="node"
      :keyword="keyword"
      :checked="checked"
      :connected="connected"
    />
  </template>
  <template v-if="type === 'label'">
    <LabelNode :node="node" :keyword="keyword" />
  </template>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { EnvironmentV1Name } from "@/components/v2";
import type { SQLEditorTreeNode as TreeNode } from "@/types";
import DatabaseNode from "./DatabaseNode.vue";
import InstanceNode from "./InstanceNode.vue";
import LabelNode from "./LabelNode.vue";

const props = defineProps<{
  node: TreeNode;
  keyword: string;
  checked: boolean;
  connected: boolean;
}>();

const type = computed(() => props.node.meta.type);
</script>
