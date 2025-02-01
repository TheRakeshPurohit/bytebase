<template>
  <div class="w-full space-y-4">
    <FeatureAttention
      v-if="!hasSensitiveDataFeature"
      feature="bb.feature.sensitive-data"
    />
    <GlobalMaskingRulesView :embedded="embedded" />
  </div>
</template>

<script lang="ts" setup>
import { FeatureAttention } from "@/components/FeatureGuard";
import { GlobalMaskingRulesView } from "@/components/SensitiveData";
import { featureToRef } from "@/store";

defineProps<{
  embedded?: boolean;
}>();

const hasSensitiveDataFeature = featureToRef("bb.feature.sensitive-data");
</script>
