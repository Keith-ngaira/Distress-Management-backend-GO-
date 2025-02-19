import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { casesAPI } from '../services/api';
import {
  Paper,
  Typography,
  Alert,
  CircularProgress,
  Box,
} from '@mui/material';
import { DataGrid } from '@mui/x-data-grid';

function CasesList() {
  const navigate = useNavigate();
  const [cases, setCases] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [page, setPage] = useState(1);
  const [pageSize, setPageSize] = useState(10);

  useEffect(() => {
    fetchCases();
  }, [page, pageSize]);

  const fetchCases = async () => {
    try {
      setLoading(true);
      setError('');
      console.log('Fetching cases...');
      const response = await casesAPI.getAllCases(page, pageSize);
      console.log('Response:', response);
      
      if (!response.data) {
        throw new Error('No data received from server');
      }

      // Ensure each case has an id and format the data
      const formattedCases = (response.data || []).map(caseItem => {
        console.log('Processing case:', caseItem);
        return {
          id: caseItem.id,
          referenceNumber: caseItem.reference_number || '',
          senderName: caseItem.sender_name || '',
          receivingDate: caseItem.receiving_date ? new Date(caseItem.receiving_date).toLocaleString('en-KE', {
            timeZone: 'Africa/Nairobi'
          }) : '',
          subject: caseItem.subject || '',
          status: caseItem.status || 'Pending',
          stage: caseItem.stage || 'Front Office Receipt'
        };
      });
      
      console.log('Formatted cases:', formattedCases);
      setCases(formattedCases);
    } catch (err) {
      console.error('Error fetching cases:', err);
      setError(err.response?.data?.message || err.message || 'Error fetching cases');
    } finally {
      setLoading(false);
    }
  };

  const columns = [
    { 
      field: 'referenceNumber', 
      headerName: 'Reference Number', 
      width: 180,
      renderCell: (params) => params.value || 'N/A'
    },
    { 
      field: 'senderName', 
      headerName: 'Sender Name', 
      width: 180,
      renderCell: (params) => params.value || 'N/A'
    },
    { 
      field: 'receivingDate', 
      headerName: 'Date Received', 
      width: 180,
      renderCell: (params) => params.value || 'N/A'
    },
    { 
      field: 'subject', 
      headerName: 'Subject', 
      width: 250,
      renderCell: (params) => params.value || 'N/A'
    },
    { 
      field: 'status', 
      headerName: 'Status', 
      width: 130,
      renderCell: (params) => params.value || 'Pending'
    },
    { 
      field: 'stage',
      headerName: 'Stage',
      width: 180,
      renderCell: (params) => params.value || 'Front Office Receipt'
    },
  ];

  if (loading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" minHeight="400px">
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Paper elevation={3} sx={{ p: 4 }}>
      <Typography variant="h5" gutterBottom>
        Distress Cases Dashboard
      </Typography>

      {error && (
        <Alert severity="error" sx={{ mb: 2 }}>
          {error}
        </Alert>
      )}
      
      <div style={{ height: 600, width: '100%' }}>
        <DataGrid
          rows={cases || []}
          columns={columns}
          pageSize={pageSize}
          page={page - 1}
          rowsPerPageOptions={[10, 25, 50]}
          onPageChange={(newPage) => setPage(newPage + 1)}
          onPageSizeChange={(newPageSize) => setPageSize(newPageSize)}
          disableSelectionOnClick
          onRowClick={(params) => navigate(`/cases/${params.id}`)}
          loading={loading}
          components={{
            NoRowsOverlay: () => (
              <Box display="flex" justifyContent="center" alignItems="center" height="100%">
                <Typography>No cases found</Typography>
              </Box>
            )
          }}
        />
      </div>
    </Paper>
  );
}

export default CasesList;
