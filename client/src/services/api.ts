import axios from 'axios';
import { Book, BookFormData } from '@/types';

const apiClient = axios.create({
  baseURL: 'http://localhost:8080', 
  headers: {
    'Content-Type': 'application/json',
  },
});

export const getBooks = async (): Promise<Book[]> => {
  const response = await apiClient.get('/books/');
  return response.data || []; 
};

export const getBookById = async (id: number): Promise<Book> => {
  const response = await apiClient.get(`/books/${id}`);
  return response.data;
};

export const createBook = async (data: BookFormData): Promise<Book> => {
  const response = await apiClient.post('/books/', data);
  return response.data;
};

export const updateBook = async (id: number, data: BookFormData): Promise<void> => {
  await apiClient.put(`/books/${id}`, data);
};

export const deleteBook = async (id: number): Promise<void> => {
  await apiClient.delete(`/books/${id}`);
};