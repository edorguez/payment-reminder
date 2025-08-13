import { RouterProvider } from 'react-router-dom';
import { router } from '../routes';
import { useAuth } from '../context/AuthContext';

const RoutesWrapper = () => {
  const { loading } = useAuth();

  if (loading) {
    return <div>Loading...</div>;
  }

  return <RouterProvider router={router} />;
};

export default RoutesWrapper;
