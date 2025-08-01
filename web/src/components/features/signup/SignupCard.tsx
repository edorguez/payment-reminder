import { useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import { useState } from 'react';
import {
    validEmail,
    validLetters,
    validLettersAndNumbers,
    validWithNoSpaces,
} from '../../../utils/InputUtils';
import type { SignUp } from '../../../types/auth';
import GoogleLogin from '../login/GoogleLogin';
import { useAuth } from '../../../context/AuthContext';
import { notify } from '../../../lib/toast';
import { FirebaseError } from 'firebase/app';
import { GetFirebaseErrorMessage } from '../../../utils/FirebaseUtils';

const SignupCard = () => {
    const { t } = useTranslation('common');
    const navigate = useNavigate();
    const { signUp } = useAuth();

    const [formData, setFormData] = useState<SignUp>({
        name: '',
        email: '',
        password: '',
        repeatPassword: '',
    });
    const [formDataErrors, setFormDataErrors] = useState<SignUp>({
        name: '',
        email: '',
        password: '',
        repeatPassword: '',
    });

    const handleNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        if (value && !validLetters(value, true)) return;
        setFormDataErrors((prev) => ({ ...prev, [name]: '' }));
        setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
    };

    const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        if (value && !validWithNoSpaces(value)) return;
        setFormDataErrors((prev) => ({ ...prev, [name]: '' }));
        setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
    };

    const handlePasswordChange = (
        event: React.ChangeEvent<HTMLInputElement>,
    ) => {
        const { name, value } = event.target;
        if (value && !validLettersAndNumbers(value)) return;
        setFormDataErrors((prev) => ({ ...prev, [name]: '' }));
        setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
    };

    const handleSignup = async (e: React.FormEvent) => {
        e.preventDefault();

        if (!isValidForm()) return;

        try {
            await signUp(
                formData.email,
                formData.password,
                formData.name
            );
            navigate('/dashboard');
        } catch (err) {
            let message = GetFirebaseErrorMessage(
                err as FirebaseError,
                t('signUp.signupError'),
            );
            notify.error(message);
        }
    };

    const isValidForm = (): boolean => {
        let isValid = true;
        setFormDataErrors({
            name: '',
            email: '',
            password: '',
            repeatPassword: '',
        });

        if (!formData.name) {
            setFormDataErrors((prev) => ({
                ...prev,
                firstName: t('signUp.firstNameEmptyError'),
            }));
            isValid = false;
        }

        if (formData.name.length < 3) {
            setFormDataErrors((prev) => ({
                ...prev,
                firstName: t('signUp.firstNameLengthError'),
            }));
            isValid = false;
        }

        if (!validEmail(formData.email)) {
            setFormDataErrors((prev) => ({
                ...prev,
                email: t('signUp.emailError'),
            }));
            isValid = false;
        }

        if (formData.password.length < 8) {
            setFormDataErrors((prev) => ({
                ...prev,
                password: t('signUp.passwordLengthError'),
            }));
            isValid = false;
        }

        if (formData.password !== formData.repeatPassword) {
            setFormDataErrors((prev) => ({
                ...prev,
                password: t('signUp.passwordMatchError'),
            }));
            isValid = false;
        }

        return isValid;
    };

    return (
        <>
            <div className="flex justify-center mt-10">
                <fieldset className="fieldset bg-base-200 border-base-300 rounded-box w-xs border p-4">
                    <label className="label">{t('signUp.firstName')}</label>
                    <input
                        type="text"
                        className={`input ${formDataErrors.name ? 'input-error' : ''}`}
                        name="name"
                        value={formData.name}
                        onChange={handleNameChange}
                        placeholder={t('signUp.firstName')}
                    />
                    <small className="text-error">
                        {formDataErrors.name}
                    </small>

                    <label className="label">{t('signUp.email')}</label>
                    <input
                        type="email"
                        className={`input ${formDataErrors.email ? 'input-error' : ''}`}
                        name="email"
                        value={formData.email}
                        onChange={handleEmailChange}
                        placeholder={t('signUp.email')}
                    />
                    <small className="text-error">{formDataErrors.email}</small>

                    <label className="label">{t('signUp.password')}</label>
                    <input
                        type="password"
                        className={`input ${formDataErrors.password ? 'input-error' : ''}`}
                        name="password"
                        value={formData.password}
                        onChange={handlePasswordChange}
                        placeholder={t('signUp.password')}
                    />
                    <small className="text-error">
                        {formDataErrors.password}
                    </small>

                    <label className="label">
                        {t('signUp.repeatPassword')}
                    </label>
                    <input
                        type="password"
                        className={`input ${formDataErrors.repeatPassword ? 'input-error' : ''}`}
                        name="repeatPassword"
                        value={formData.repeatPassword}
                        onChange={handlePasswordChange}
                        placeholder={t('signUp.repeatPassword')}
                    />
                    <small className="text-error">
                        {formDataErrors.repeatPassword}
                    </small>

                    <button
                        className="btn btn-neutral mt-4"
                        onClick={handleSignup}
                    >
                        {t('signUp.signUp')}
                    </button>

                    <div className="flex items-center my-2">
                        <div className="w-full h-[1px] bg-gray-400"></div>
                        <span className="mx-2">{t('login.or')}</span>
                        <div className="w-full h-[1px] bg-gray-400"></div>
                    </div>

                    <div className="flex justify-center">
                        <GoogleLogin />
                    </div>
                </fieldset>
            </div>
        </>
    );
};

export default SignupCard;
