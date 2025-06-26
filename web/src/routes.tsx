import { createBrowserRouter, Navigate } from 'react-router-dom';
import Home from './pages/Home';
import NotFound from './pages/NotFound';
import PublicLayout from './layouts/PublicLayout';
import PrivateLayout from './layouts/PrivateLayout';
import Dashboard from './pages/Dashboard';
import Login from './pages/Login';
import { useAuth } from './context/AuthContext';
import Signup from './pages/Signup';

const PrivateRoute = () => {
  const { isAuthenticated } = useAuth();
  return isAuthenticated ? <PrivateLayout /> : <Navigate to="/login" replace />;
};

export const router = createBrowserRouter([
  {
    element: <PublicLayout />,
    children: [
      { index: true, element: <Home /> },
      { path: '/login', element: <Login /> },
      { path: '/signup', element: <Signup /> },
    ]
  },
  {
    element: <PrivateRoute />,
    children: [
      { path: '/dashboard', element: <Dashboard /> },
    ]
  },
  {
    path: '*',
    element: <NotFound />
  }

]);
