import { useNavigate } from "react-router-dom";
import { useTranslation } from "react-i18next";
import { useState } from "react";
import { validEmail, validLetters, validLettersAndNumbers, validWithNoSpaces } from "../../../utils/InputUtils";
import type { SignUp } from "../../../types/auth";
import GoogleLogin from "../login/GoogleLogin";
import { useAuth } from "../../../context/AuthContext";
import { notify } from "../../../lib/toast";
import { FirebaseError } from "firebase/app";

const SignupCard = () => {
  const { t } = useTranslation('common');
  const navigate = useNavigate();
  const { signUp } = useAuth();

  const [formData, setFormData] = useState<SignUp>({ firstName: "", lastName: "", email: "", password: "", repeatPassword: "" });
  const [formDataErrors, setFormDataErrors] = useState<SignUp>({ firstName: "", lastName: "", email: "", password: "", repeatPassword: "" });

  const handleNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    if (value && !validLetters(value, true)) return;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

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

  const handleSignup = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!isValidForm())
      return;

    try {
      await signUp(formData.email, formData.password, formData.firstName, formData.lastName);
      navigate('/dashboard');
    } catch (err) {
      const message = err instanceof FirebaseError ? err.message : 'Unknown error';
      notify.error(message || 'Login failed');
    }
  };

  const isValidForm = (): boolean => {
    let isValid = true;
    setFormDataErrors({ firstName: "", lastName: "", email: "", password: "", repeatPassword: "" });

    if (!formData.firstName) {
      setFormDataErrors((prev) => ({ ...prev, firstName: t('signUp.firstNameEmptyError') }));
      isValid = false;
    }

    if (!formData.lastName) {
      setFormDataErrors((prev) => ({ ...prev, lastName: t('signUp.lastNameEmptyError') }));
      isValid = false;
    }

    if (formData.firstName.length < 3) {
      setFormDataErrors((prev) => ({ ...prev, firstName: t('signUp.firstNameLengthError') }));
      isValid = false;
    }

    if (formData.lastName.length < 3) {
      setFormDataErrors((prev) => ({ ...prev, lastName: t('signUp.lastNameLengthError') }));
      isValid = false;
    }

    if (!validEmail(formData.email)) {
      setFormDataErrors((prev) => ({ ...prev, email: t('signUp.emailError') }));
      isValid = false;
    }

    if (formData.password.length < 8) {
      setFormDataErrors((prev) => ({ ...prev, password: t('signUp.passwordLengthError') }));
      isValid = false;
    }

    if (formData.password !== formData.repeatPassword) {
      setFormDataErrors((prev) => ({ ...prev, password: t('signUp.passwordMatchError') }));
      isValid = false;
    }

    return isValid;
  }

  return (
    <>
      <div className="flex justify-center mt-10">
        <fieldset className="fieldset bg-base-200 border-base-300 rounded-box w-xs border p-4">

          <label className="label">{t('signUp.firstName')}</label>
          <input type="text" className={`input ${formDataErrors.firstName ? 'input-error' : ''}`} name="firstName" value={formData.firstName} onChange={handleNameChange} placeholder={t('signUp.firstName')} />
          <small className="text-error">{formDataErrors.firstName}</small>

          <label className="label">{t('signUp.lastName')}</label>
          <input type="text" className={`input ${formDataErrors.lastName ? 'input-error' : ''}`} name="lastName" value={formData.lastName} onChange={handleNameChange} placeholder={t('signUp.lastName')} />
          <small className="text-error">{formDataErrors.lastName}</small>

          <label className="label">{t('signUp.email')}</label>
          <input type="email" className={`input ${formDataErrors.email ? 'input-error' : ''}`} name="email" value={formData.email} onChange={handleEmailChange} placeholder={t('signUp.email')} />
          <small className="text-error">{formDataErrors.email}</small>

          <label className="label">{t('signUp.password')}</label>
          <input type="password" className={`input ${formDataErrors.password ? 'input-error' : ''}`} name="password" value={formData.password} onChange={handlePasswordChange} placeholder={t('signUp.password')} />
          <small className="text-error">{formDataErrors.password}</small>

          <label className="label">{t('signUp.repeatPassword')}</label>
          <input type="password" className={`input ${formDataErrors.repeatPassword ? 'input-error' : ''}`} name="repeatPassword" value={formData.repeatPassword} onChange={handlePasswordChange} placeholder={t('signUp.repeatPassword')} />
          <small className="text-error">{formDataErrors.repeatPassword}</small>

          <button className="btn btn-neutral mt-4" onClick={handleSignup}>{t('signUp.signUp')}</button>

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

export default SignupCard;
