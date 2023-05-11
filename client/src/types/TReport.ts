export type TReport = {
  id: string;
  external_id: string;
  image_url: string;
  priority: number;
  status: number;
  created_at: string;
};

export type CreateReportDTO = {
  id?: string;
  external_id: string;
  image_url: string;
  priority?: number;
  status?: number;
};
