import React from 'react';
import { auth, googleProvider, signInWithPopup } from '../../../firebase';
import { FirebaseError } from 'firebase/app';
import { useNavigate } from 'react-router-dom';
import { FcGoogle } from "react-icons/fc";
import { getIdToken } from 'firebase/auth';

const GoogleLogin: React.FC = () => {
  const navigate = useNavigate();

  const signIn = async () => {
    try {
      const userCredential = await signInWithPopup(auth, googleProvider);

      const user = userCredential.user;
      const firebaseUid = user.uid;
      const email = user.email;            // may be null if account has no email
      const fullName = user.displayName;   // may be null

      const idToken = await getIdToken(user);

      console.log({ firebaseUid, email, fullName, idToken });

      navigate('/dashboard');
    } catch (err) {
      const message = err instanceof FirebaseError ? err.message : 'Unknown error';
      console.log('hola')
      console.log(message);
    }
  };

  return (
    <button onClick={signIn} className='text-2xl rounded-full bg-white p-1 cursor-pointer transition-transform duration-200 hover:scale-120'>
      <FcGoogle />
    </button>
  );
};

export default GoogleLogin;
