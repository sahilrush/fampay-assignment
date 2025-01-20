export interface TGetVideo {
  id: number;
  title: string;
  description: string;
  publishedAt: string;
  thumbnails: string;
}

export interface TGetVideosResponse {
  limit: number;
  page: number;
  videos: TGetVideo[];
}