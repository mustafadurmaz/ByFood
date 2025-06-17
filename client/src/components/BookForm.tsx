"use client";

import { useForm, Controller } from "react-hook-form";
import { Book, BookFormData } from "@/types";
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Button,
  Box,
} from "@mui/material";

interface BookFormProps {
  open: boolean;
  onClose: () => void;
  onSubmit: (data: BookFormData) => void;
  defaultValues?: Book;
}

export const BookForm = ({
  open,
  onClose,
  onSubmit,
  defaultValues,
}: BookFormProps) => {
  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<BookFormData>({
    defaultValues: {
      title: defaultValues?.title || "",
      author: defaultValues?.author || "",
      year: defaultValues?.year || new Date().getFullYear(),
    },
  });

  return (
    <Dialog open={open} onClose={onClose} fullWidth maxWidth="sm">
      <DialogTitle>{defaultValues ? "Edit Book" : "Add Book"}</DialogTitle>
      <form onSubmit={handleSubmit(onSubmit)}>
        <DialogContent>
          <Box sx={{ display: "flex", flexDirection: "column", gap: 2, pt: 1 }}>
            <Controller
              name="title"
              control={control}
              rules={{ required: "This field is required" }}
              render={({ field }) => (
                <TextField
                  {...field}
                  label="Title"
                  fullWidth
                  error={!!errors.title}
                  helperText={errors.title?.message}
                />
              )}
            />
            <Controller
              name="author"
              control={control}
              rules={{ required: "This field is required" }}
              render={({ field }) => (
                <TextField
                  {...field}
                  label="Author"
                  fullWidth
                  error={!!errors.author}
                  helperText={errors.author?.message}
                />
              )}
            />

            <Controller
              name="year"
              control={control}
              rules={{
                required: "This field is required",
                validate: (value) =>
                  (typeof value === "number" && !isNaN(value) && value > 0) ||
                  "Please enter a valid year",
              }}
              render={({ field }) => (
                <TextField
                  {...field}
                  label="Publish Year"
                  type="number"
                  fullWidth
                  error={!!errors.year}
                  helperText={errors.year?.message}
                  onChange={(e) => {
                    const value = e.target.value;
                    field.onChange(value === "" ? "" : parseInt(value, 10));
                  }}
                />
              )}
            />
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={onClose}>Cancel</Button>
          <Button type="submit" variant="contained">
            Save
          </Button>
        </DialogActions>
      </form>
    </Dialog>
  );
};
