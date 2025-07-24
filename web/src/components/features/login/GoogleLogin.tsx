import React from 'react';
import { FirebaseError } from 'firebase/app';
import { useNavigate } from 'react-router-dom';
import { FcGoogle } from "react-icons/fc";
import { useAuth } from '../../../context/AuthContext';

const GoogleLogin: React.FC = () => {
  const navigate = useNavigate();
  const { loginGoogle } = useAuth();

  const login = async () => {
    try {
      await loginGoogle();
      navigate('/dashboard');
    } catch (err) {
      const message = err instanceof FirebaseError ? err.message : 'Unknown error';
      console.log('hola')
      console.log(message);
    }
  };

  return (
    <button onClick={login} className='text-2xl rounded-full bg-white p-1 cursor-pointer transition-transform duration-200 hover:scale-120'>
      <FcGoogle />
    </button>
  );
};

export default GoogleLogin;
