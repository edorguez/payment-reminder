import React from 'react';
import { auth, googleProvider, signInWithPopup } from '../../../firebase';
import { FirebaseError } from 'firebase/app';
import { useNavigate } from 'react-router-dom';
import { FcGoogle } from "react-icons/fc";

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
  return (
    <button onClick={signIn} className='text-2xl rounded-full bg-white p-1 cursor-pointer transition-transform duration-200 hover:scale-120'>
      <FcGoogle />
    </button>
  );
};

export default GoogleLogin;
