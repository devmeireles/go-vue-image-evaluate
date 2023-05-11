import type { CreateReportDTO, TReport } from "@/types/TReport";
import axios from "axios";
import { defineStore } from "pinia";

interface ReportState {
  reports: TReport[];
}

export const useReportStore = defineStore({
  id: "report",
  state: () =>
    ({
      reports: []
    } as ReportState),
  getters: {
    getReportById: (state) => (id: string) => {
      return state.reports.find((reports: TReport) => reports.id === id);
    },
    getReports: (state) => () => state.reports
  },
  actions: {
    async fetchReport(id: string) {
      try {
        const response = await axios.get(`http://localhost:3000/report/${id}`);

        const { data, status } = response;

        if (data.success && status === 200) {
          return data.data;
        }
      } catch (error) {
        console.log(error);
      }
    },
    async fetchReports() {
      try {
        const response = await axios.get("http://localhost:3000/report");

        const { data, status } = response;

        if (data.success && status === 200) {
          this.$patch({
            reports: data.data
          });
        }
      } catch (error) {
        console.log(error);
      }
    },

    async saveReport(payload: CreateReportDTO) {
      try {
        const response = await axios.post(
          "http://localhost:3000/report",
          payload
        );

        const { data, status } = response;

        if (status === 201 && data.success) {
          return data.data;
        }
      } catch (error) {
        console.log(error);
      }
    },

    async patchReport(payload: CreateReportDTO) {
      try {
        const response = await axios.patch(
          `http://localhost:3000/report/${payload.id}`,
          payload
        );

        const { data, status } = response;

        if (status === 200 && data.success) {
          return data.data;
        }
      } catch (error) {
        console.log(error);
      }
    }
  }
});
