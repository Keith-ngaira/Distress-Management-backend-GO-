import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { casesAPI } from '../services/api';
import {
  Container,
  Paper,
  TextField,
  Typography,
  Button,
  Grid,
  Alert,
  MenuItem,
  Box,
  CircularProgress
} from '@mui/material';
import UploadFileIcon from '@mui/icons-material/UploadFile';

const DistressForm = () => {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    sender_name: '',
    subject: '',
    country_of_origin: '',
    distressed_person_name: '',
    nature_of_case: '',
    case_details: '',
    receiving_date: new Date().toISOString()
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [files, setFiles] = useState([]);
  const [uploading, setUploading] = useState(false);
  const [uploadError, setUploadError] = useState('');

  const natureOfCases = [
    'Emergency',
    'Urgent',
    'Standard'
  ];

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prevState => ({
      ...prevState,
      [name]: value
    }));
  };

  const handleFileChange = (event) => {
    const selectedFiles = Array.from(event.target.files);
    setFiles(selectedFiles);
  };

  const uploadFiles = async (caseId) => {
    setUploading(true);
    setUploadError('');
    
    try {
      for (const file of files) {
        const formData = new FormData();
        formData.append('document', file);
        
        await casesAPI.uploadDocument(caseId, formData);
      }
    } catch (error) {
      console.error('Error uploading files:', error);
      setUploadError('Error uploading files. Please try again.');
      throw error;
    } finally {
      setUploading(false);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      // Create the case first
      const response = await casesAPI.createCase({
        sender_name: formData.sender_name,
        subject: formData.subject,
        country_of_origin: formData.country_of_origin,
        distressed_person_name: formData.distressed_person_name,
        nature_of_case: formData.nature_of_case,
        case_details: formData.case_details,
        status: 'Pending',
        stage: 'Front Office Receipt'
      });

      // If case creation successful, upload files
      if (files.length > 0) {
        await uploadFiles(response.data.id);
      }

      navigate(`/cases/${response.data.id}`);
    } catch (err) {
      console.error('Error:', err);
      setError(err.response?.data?.message || 'Error creating case');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Container maxWidth="md" sx={{ mt: 4 }}>
      <Paper elevation={3} sx={{ p: 4 }}>
        <Typography variant="h5" gutterBottom>
          New Distress Case
        </Typography>

        {error && (
          <Alert severity="error" sx={{ mb: 2 }}>
            {error}
          </Alert>
        )}

        {uploadError && (
          <Alert severity="error" sx={{ mb: 2 }}>
            {uploadError}
          </Alert>
        )}

        <form onSubmit={handleSubmit}>
          <Grid container spacing={3}>
            <Grid item xs={12} sm={6}>
              <TextField
                required
                fullWidth
                id="sender_name"
                label="Sender Name"
                name="sender_name"
                value={formData.sender_name}
                onChange={handleChange}
                inputProps={{ 'aria-label': 'sender name' }}
              />
            </Grid>

            <Grid item xs={12} sm={6}>
              <TextField
                required
                fullWidth
                id="subject"
                label="Subject"
                name="subject"
                value={formData.subject}
                onChange={handleChange}
                inputProps={{ 'aria-label': 'subject' }}
              />
            </Grid>

            <Grid item xs={12} sm={6}>
              <TextField
                required
                fullWidth
                id="country_of_origin"
                label="Country of Origin"
                name="country_of_origin"
                value={formData.country_of_origin}
                onChange={handleChange}
                inputProps={{ 'aria-label': 'country of origin' }}
              />
            </Grid>

            <Grid item xs={12} sm={6}>
              <TextField
                required
                fullWidth
                id="distressed_person_name"
                label="Distressed Person Name"
                name="distressed_person_name"
                value={formData.distressed_person_name}
                onChange={handleChange}
                inputProps={{ 'aria-label': 'distressed person name' }}
              />
            </Grid>

            <Grid item xs={12}>
              <TextField
                required
                fullWidth
                select
                id="nature_of_case"
                label="Nature of Case"
                name="nature_of_case"
                value={formData.nature_of_case}
                onChange={handleChange}
                inputProps={{ 'aria-label': 'nature of case' }}
              >
                {natureOfCases.map((option) => (
                  <MenuItem key={option} value={option}>
                    {option}
                  </MenuItem>
                ))}
              </TextField>
            </Grid>

            <Grid item xs={12}>
              <TextField
                required
                fullWidth
                multiline
                rows={4}
                id="case_details"
                label="Case Details"
                name="case_details"
                value={formData.case_details}
                onChange={handleChange}
                inputProps={{ 'aria-label': 'case details' }}
              />
            </Grid>

            <Box sx={{ mb: 2 }}>
              <Typography variant="subtitle1" gutterBottom>
                Supporting Documents
              </Typography>
              <input
                type="file"
                multiple
                onChange={handleFileChange}
                accept=".pdf,.doc,.docx,.xls,.xlsx,.jpg,.jpeg,.png,.gif"
                style={{ display: 'none' }}
                id="document-upload"
              />
              <label htmlFor="document-upload">
                <Button
                  variant="outlined"
                  component="span"
                  startIcon={<UploadFileIcon />}
                  sx={{ mr: 2 }}
                >
                  Select Files
                </Button>
              </label>
              {files.length > 0 && (
                <Typography variant="body2" color="textSecondary">
                  {files.length} file(s) selected
                </Typography>
              )}
            </Box>

            <Grid item xs={12}>
              <Button
                type="submit"
                variant="contained"
                color="primary"
                fullWidth
                disabled={loading || uploading}
                startIcon={loading || uploading ? <CircularProgress size={20} /> : null}
              >
                {loading || uploading ? 'Submitting...' : 'Submit Case'}
              </Button>
            </Grid>
          </Grid>
        </form>
      </Paper>
    </Container>
  );
};

export default DistressForm;
