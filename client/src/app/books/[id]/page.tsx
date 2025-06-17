"use client";

import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import { Book } from "@/types";
import { getBookById } from "@/services/api";
import {
  Container,
  Card,
  CardContent,
  Typography,
  CircularProgress,
  Box,
  Alert,
} from "@mui/material";

export default function BookDetailPage() {
  const params = useParams();
  const id = Number(params.id);
  const [book, setBook] = useState<Book | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (id) {
      setLoading(true);
      getBookById(id)
        .then((data) => {
          setBook(data);
          setError(null);
        })
        .catch(() => {
          setError("Book not found");
        })
        .finally(() => {
          setLoading(false);
        });
    }
  }, [id]);

  if (loading) {
    return (
      <Container>
        <Box sx={{ display: "flex", justifyContent: "center", mt: 5 }}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  if (error) {
    return (
      <Container>
        <Alert severity="error" sx={{ mt: 5 }}>
          {error}
        </Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="lg" sx={{ mt: 4 }}>
      {book && (
        <Card>
          <CardContent>
            <Typography variant="h3" component="h1" gutterBottom>
              {book.title}
            </Typography>
            <Typography variant="h5" color="text.secondary">
              by {book.author}
            </Typography>
            <Typography sx={{ mt: 2 }}>Publish Year: {book.year}</Typography>
            <Typography variant="body2" color="text.secondary" sx={{ mt: 1 }}>
              Added Date: {new Date(book.created_at).toLocaleDateString()}
            </Typography>
          </CardContent>
        </Card>
      )}
    </Container>
  );
}
