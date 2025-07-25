import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { RouterProvider } from 'react-router-dom';
import './index.css';
import './i18n/i18n';
import { router } from './routes';
import { AuthProvider } from './context/AuthContext';
import { ToastContainer } from './components/ui/ToastContainer';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <AuthProvider>
      <RouterProvider router={router} /> 
    </AuthProvider>
    <ToastContainer />
  </StrictMode>,
)
