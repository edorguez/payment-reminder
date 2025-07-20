import React from 'react';
import { auth, googleProvider, signInWithPopup } from '../../../firebase';
import { FirebaseError } from 'firebase/app';
import { useNavigate } from 'react-router-dom';

const GoogleLogin: React.FC = () => {
  const navigate = useNavigate();

  const signIn = async () => {
    try {
      await signInWithPopup(auth, googleProvider);
      navigate('/dashboard');
    } catch (err) {
      const message = err instanceof FirebaseError ? err.message : 'Unknown error';
      alert(message);
    }
  };
  return <button onClick={signIn}>Sign in with Google</button>;
};

export default GoogleLogin;
