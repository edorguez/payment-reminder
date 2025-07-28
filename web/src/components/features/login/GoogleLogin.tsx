import React from 'react';
import { FirebaseError } from 'firebase/app';
import { useNavigate } from 'react-router-dom';
import { FcGoogle } from 'react-icons/fc';
import { useAuth } from '../../../context/AuthContext';
import { GetFirebaseErrorMessage } from '../../../utils/FirebaseUtils';
import { notify } from '../../../lib/toast';
import { useTranslation } from 'react-i18next';

const GoogleLogin: React.FC = () => {
    const { t } = useTranslation('common');
    const navigate = useNavigate();
    const { loginGoogle } = useAuth();

    const login = async () => {
        try {
            await loginGoogle();
            navigate('/dashboard');
        } catch (err) {
            let message = GetFirebaseErrorMessage(
                err as FirebaseError,
                t('login.loginError'),
            );
            notify.error(message);
        }
    };

    return (
        <button
            onClick={login}
            className="text-2xl rounded-full bg-white p-1 cursor-pointer transition-transform duration-200 hover:scale-120"
        >
            <FcGoogle />
        </button>
    );
};

export default GoogleLogin;
