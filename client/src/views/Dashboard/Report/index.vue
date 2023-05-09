<template>
  <CustomCard :title="$t('report.title_plural')" :createRoute="createRoute">
    <ReportsList :data="reports" />
  </CustomCard>
</template>

<script lang="ts" setup>
import CustomCard from "@/components/organisms/CustomCard.vue";
import ReportsList from "@/components/organisms/ReportsList.vue";
import { routes } from "@/consts/routes";
import { useReportStore } from "@/stores/report";
import { computed, onMounted } from "vue";

const createRoute = routes.report.create;
const reportStore = useReportStore();

onMounted(() => {
  reportStore.fetchReports();
});

const reports = computed(() => {
  return reportStore.getReports();
});
</script>
