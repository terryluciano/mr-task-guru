export interface Task {
  id: number;
  name: string;
  category?: number | string;
  minutes?: number;
  current_status?: string;
  timestamp: string | Date;
}
