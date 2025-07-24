import { useNavigate } from "react-router-dom";
import { useTranslation } from "react-i18next";
import { useState } from "react";
import { validLetters, validLettersAndNumbers, validWithNoSpaces } from "../../../utils/InputUtils";
import type { SignUp } from "../../../types/auth";
import GoogleLogin from "../login/GoogleLogin";
import { useAuth } from "../../../context/AuthContext";

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

    if(!isValidForm())
      return;

    try {
      await signUp(formData.email, formData.password, formData.firstName, formData.lastName);
      navigate('/dashboard');
    } catch (err) {
      console.log('Error normal singup')
      console.log(err)
    }
  };

  const isValidForm = (): boolean => {
    if(!formData.firstName) {
      setFormDataErrors((prev) => ({ ...prev, firstName: t('signUp.firstNameEmptyError')}));
      return false;
    }

    if(formData.firstName.length < 3 || formData.lastName.length < 3)
      return false;

    if (formData.password !== formData.repeatPassword) {
      console.log('password doesnt match')
      return false;
    }

    return true;
  }

  return (
    <>
      <div className="flex justify-center mt-20">
        <fieldset className="fieldset bg-base-200 border-base-300 rounded-box w-xs border p-4">

          <label className="label">{t('signUp.firstName')}</label>
          <input type="text" className="input" name="firstName" value={formData.firstName} onChange={handleNameChange} placeholder={t('signUp.firstName')} />

          <label className="label">{t('signUp.lastName')}</label>
          <input type="text" className="input" name="lastName" value={formData.lastName} onChange={handleNameChange} placeholder={t('signUp.lastName')} />

          <label className="label">{t('signUp.email')}</label>
          <input type="email" className="input" name="email" value={formData.email} onChange={handleEmailChange} placeholder={t('signUp.email')} />

          <label className="label">{t('signUp.password')}</label>
          <input type="password" className="input" name="password" value={formData.password} onChange={handlePasswordChange} placeholder={t('signUp.password')} />

          <label className="label">{t('signUp.repeatPassword')}</label>
          <input type="password" className="input" name="repeatPassword" value={formData.repeatPassword} onChange={handlePasswordChange} placeholder={t('signUp.repeatPassword')} />

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
