'use client';

import { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { Book, BookFormData } from '@/types';
import * as api from '@/services/api';
import toast from 'react-hot-toast';

interface BookContextType {
  books: Book[];
  loading: boolean;
  fetchBooks: () => void;
  addBook: (book: BookFormData) => Promise<void>;
  editBook: (id: number, book: BookFormData) => Promise<void>;
  removeBook: (id: number) => Promise<void>;
}

const BookContext = createContext<BookContextType | undefined>(undefined);

export const BookProvider = ({ children }: { children: ReactNode }) => {
  const [books, setBooks] = useState<Book[]>([]);
  const [loading, setLoading] = useState(true);

  const fetchBooks = async () => {
    setLoading(true);
    try {
      const data = await api.getBooks();
      setBooks(data);
    } catch (error) {
      toast.error('An error occurred while loading books.');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchBooks();
  }, []);

  const addBook = async (bookData: BookFormData) => {
    await toast.promise(
      api.createBook(bookData),
      {
        loading: 'Book adding...',
        success: 'Book added successfully',
        error: 'An error occurred while adding the book.',
      }
    );
    fetchBooks(); 
  };
  
  const editBook = async (id: number, bookData: BookFormData) => {
    await toast.promise(
      api.updateBook(id, bookData),
      {
        loading: 'Book updating...',
        success: 'Book updated successfully',
        error: 'An error occurred while updating the book.',
      }
    );
    fetchBooks();
  };

  const removeBook = async (id: number) => {
    await toast.promise(
      api.deleteBook(id),
      {
        loading: 'Book deleting...',
        success: 'Book deleted successfully',
        error: 'An error occurred while deleting the book.',
      }
    );
    fetchBooks();
  };

  return (
    <BookContext.Provider value={{ books, loading, fetchBooks, addBook, editBook, removeBook }}>
      {children}
    </BookContext.Provider>
  );
};

export const useBooks = () => {
  const context = useContext(BookContext);
  if (context === undefined) {
    throw new Error('useBooks must be used within a BookProvider');
  }
  return context;
};