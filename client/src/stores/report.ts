import type { CreateReportDTO, TReport } from "@/types/TReport";
import axios from "axios";
import { defineStore } from "pinia";

interface ReportState {
  reports: TReport[];
}

const BASE_URL = import.meta.env.VITE_API_BASE_URL;

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
        const response = await axios.get(`${BASE_URL}/report/${id}`);

        const { data, status } = response;

        if (data.success && status === 200) {
          return data.data;
        }
      } catch (error) {
        console.log(error);
      }
    },
    async fetchReports(filters?: any[]) {
      try {
        let params = {};

        if (filters && filters.length > 0) {
          filters.map((item) => {
            Object.assign(params, item);
          });
        }

        const response = await axios.get(`${BASE_URL}/report`, {
          params
        });

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
        const response = await axios.post(`${BASE_URL}/report`, payload);

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
          `${BASE_URL}/report/${payload.id}`,
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
