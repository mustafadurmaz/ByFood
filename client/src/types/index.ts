export interface Book {
  id: number;
  title: string;
  author: string;
  year: number;
  created_at: string;
}

export type BookFormData = Omit<Book, 'id' | 'created_at'>;