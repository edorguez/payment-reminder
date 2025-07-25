import { Toaster } from 'react-hot-toast';

export const ToastContainer = () => (
  <Toaster
    position="top-right"
    toastOptions={{
      duration: 4000,
      style: {},                 
      className: 'alert',       
      success: {
        className: 'alert alert-success',
      },
      error: {
        className: 'alert alert-error',
      },
      loading: {
        className: 'alert alert-info',
      },
    }}
  />
);
