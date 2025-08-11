import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import './index.css';
import './i18n/i18n';
import { AuthProvider } from './context/AuthContext';
import { ToastContainer } from './components/ui/ToastContainer';
import { ProfileProvider } from './context/ProfileContext';
import RoutesWrapper from './components/RoutesWrapper';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <AuthProvider>
      <ProfileProvider>
        <RoutesWrapper />
      </ProfileProvider>
    </AuthProvider>
    <ToastContainer />
  </StrictMode>,
)
