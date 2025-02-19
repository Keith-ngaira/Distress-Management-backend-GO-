import axios from 'axios';

const baseURL = 'http://localhost:8080/api';

const api = axios.create({
  baseURL,
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true
});

// Request interceptor
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

// Auth APIs
export const authAPI = {
  login: async (credentials) => {
    const response = await api.post('/auth/login', credentials);
    if (response.data.token) {
      localStorage.setItem('token', response.data.token);
    }
    return response;
  },
  register: (userData) => api.post('/auth/register', userData),
  logout: () => {
    localStorage.removeItem('token');
    window.location.href = '/login';
  }
};

// Cases APIs
export const casesAPI = {
  getAllCases: (page = 1, limit = 10) => api.get(`/cases?page=${page}&limit=${limit}`),
  getCase: (id) => api.get(`/cases/${id}`),
  createCase: (caseData) => api.post('/cases', caseData),
  updateCase: (id, caseData) => api.put(`/cases/${id}`, caseData),
  deleteCase: (id) => api.delete(`/cases/${id}`),
  updateStatus: (id, status) => api.patch(`/cases/${id}/status`, { status }),
  addProgressNote: (id, note) => api.post(`/cases/${id}/progress-notes`, note),
  // Document endpoints
  uploadDocument: (caseId, formData) => api.post(`/cases/${caseId}/documents`, formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  }),
  getDocuments: (caseId) => api.get(`/cases/${caseId}/documents`),
  deleteDocument: (caseId, documentId) => api.delete(`/cases/${caseId}/documents/${documentId}`)
};

// Users APIs
export const usersAPI = {
  getAllUsers: () => api.get('/users'),
  getUser: (id) => api.get(`/users/${id}`),
  updateUser: (id, userData) => api.put(`/users/${id}`, userData),
  deleteUser: (id) => api.delete(`/users/${id}`)
};

export default api;
