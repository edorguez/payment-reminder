import { useNavigate } from "react-router-dom";
import { useAuth } from "../../../context/AuthContext";
import { useTranslation } from "react-i18next";
import GoogleLogin from "./GoogleLogin";
import { useState } from "react";
import type { Login } from "../../../types/auth";
import { validEmail, validLettersAndNumbers, validWithNoSpaces } from "../../../utils/InputUtils";
import { FirebaseError } from "firebase/app";
import { notify } from "../../../lib/toast";

const LoginCard = () => {
  const { t } = useTranslation('common');
  const navigate = useNavigate();
  const { login } = useAuth();
  const [formData, setFormData] = useState<Login>({ email: "", password: "" });
  const [formDataErrors, setFormDataErrors] = useState<Login>({ email: "", password: "" });

  const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    if (value && !validWithNoSpaces(value)) return;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

  const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    if (value && !validLettersAndNumbers(value)) return;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

  const handleLogin = async (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    
    if(!isValidForm())
      return;

    try {
      await login(formData.email, formData.password);
      navigate('/dashboard');
    } catch (err) {
      console.log(err)
      const message = err instanceof FirebaseError ? err.message : 'Unknown error';
      notify.error(message || 'Login failed');
    }
  };

  const isValidForm = (): boolean => {
    let isValid = true;
    setFormDataErrors({ email: "", password: "" });

    if (!validEmail(formData.email)) {
      setFormDataErrors((prev) => ({ ...prev, email: t('login.emailError') }));
      isValid = false;
    }

    if (formData.password.length < 8) {
      setFormDataErrors((prev) => ({ ...prev, password: t('login.passwordLengthError') }));
      isValid = false;
    }

    return isValid;
  }

  return (
    <>
      <div className="flex justify-center mt-10">
        <fieldset className="fieldset bg-base-200 border-base-300 rounded-box w-xs border p-4">
          <label className="label">{t('login.email')}</label>
          <input type="email" name="email" className={`input ${formDataErrors.email ? 'input-error' : ''}`} value={formData.email} onChange={handleEmailChange} placeholder={t('login.email')} />
          <small className="text-error">{formDataErrors.email}</small>

          <label className="label">{t('login.password')}</label>
          <input type="password" name="password" className={`input ${formDataErrors.password ? 'input-error' : ''}`} value={formData.password} onChange={handlePasswordChange} placeholder={t('login.password')} />
          <small className="text-error">{formDataErrors.password}</small>

          <button className="btn btn-neutral mt-4" onClick={handleLogin}>{t('login.login')}</button>

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
}

export default LoginCard;
