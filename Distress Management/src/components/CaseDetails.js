import React, { useState, useEffect, useCallback } from 'react';
import { useParams } from 'react-router-dom';
import { casesAPI } from '../services/api';
import {
  Container,
  Paper,
  Typography,
  TextField,
  Button,
  Box,
  CircularProgress,
  Alert,
  Divider,
  List,
  ListItem,
  ListItemText
} from '@mui/material';

const CaseDetails = () => {
  const { id } = useParams();
  const [caseData, setCaseData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [note, setNote] = useState('');
  const [statusUpdate, setStatusUpdate] = useState('');

  const fetchCaseDetails = useCallback(async () => {
    try {
      const response = await casesAPI.getCase(id);
      setCaseData(response.data);
    } catch (err) {
      setError(err.response?.data?.message || 'Failed to fetch case details');
    } finally {
      setLoading(false);
    }
  }, [id]);

  useEffect(() => {
    fetchCaseDetails();
  }, [fetchCaseDetails]);

  const handleAddNote = async (e) => {
    e.preventDefault();
    if (!note.trim()) return;

    try {
      await casesAPI.addProgressNote(id, { note: note.trim() });
      setNote('');
      fetchCaseDetails(); // Refresh case details to show new note
    } catch (err) {
      setError(err.response?.data?.message || 'Failed to add note');
    }
  };

  const handleUpdateStatus = async (e) => {
    e.preventDefault();
    if (!statusUpdate.trim()) return;

    try {
      await casesAPI.updateStatus(id, statusUpdate.trim());
      setStatusUpdate('');
      fetchCaseDetails(); // Refresh case details to show new status
    } catch (err) {
      setError(err.response?.data?.message || 'Failed to update status');
    }
  };

  if (loading) {
    return (
      <Container sx={{ display: 'flex', justifyContent: 'center', mt: 4 }}>
        <CircularProgress />
      </Container>
    );
  }

  if (error) {
    return (
      <Container sx={{ mt: 4 }}>
        <Alert severity="error">{error}</Alert>
      </Container>
    );
  }

  if (!caseData) {
    return (
      <Container sx={{ mt: 4 }}>
        <Alert severity="info">Case not found</Alert>
      </Container>
    );
  }

  return (
    <Container sx={{ mt: 4 }}>
      <Paper sx={{ p: 3 }}>
        <Typography variant="h4" gutterBottom>
          Case Details
        </Typography>

        <Box sx={{ mb: 3 }}>
          <Typography variant="subtitle1" color="textSecondary">
            Reference Number: {caseData.reference_number}
          </Typography>
          <Typography variant="subtitle1" color="textSecondary">
            Status: {caseData.status}
          </Typography>
          <Typography variant="body1" sx={{ mt: 2 }}>
            {caseData.description}
          </Typography>
        </Box>

        <Divider sx={{ my: 3 }} />

        {/* Status Update Form */}
        <Box component="form" onSubmit={handleUpdateStatus} sx={{ mb: 4 }}>
          <Typography variant="h6" gutterBottom>
            Update Status
          </Typography>
          <Box sx={{ display: 'flex', gap: 2 }}>
            <TextField
              fullWidth
              label="New Status"
              value={statusUpdate}
              onChange={(e) => setStatusUpdate(e.target.value)}
              size="small"
            />
            <Button
              type="submit"
              variant="contained"
              disabled={!statusUpdate.trim()}
            >
              Update
            </Button>
          </Box>
        </Box>

        {/* Progress Notes */}
        <Typography variant="h6" gutterBottom>
          Progress Notes
        </Typography>
        
        <Box component="form" onSubmit={handleAddNote} sx={{ mb: 3 }}>
          <Box sx={{ display: 'flex', gap: 2 }}>
            <TextField
              fullWidth
              label="Add Note"
              value={note}
              onChange={(e) => setNote(e.target.value)}
              size="small"
            />
            <Button
              type="submit"
              variant="contained"
              disabled={!note.trim()}
            >
              Add
            </Button>
          </Box>
        </Box>

        <List>
          {caseData.progress_notes?.map((note, index) => (
            <ListItem key={index} divider={index < caseData.progress_notes.length - 1}>
              <ListItemText
                primary={note.note}
                secondary={new Date(note.created_at).toLocaleString()}
              />
            </ListItem>
          ))}
        </List>
      </Paper>
    </Container>
  );
};

export default CaseDetails;
