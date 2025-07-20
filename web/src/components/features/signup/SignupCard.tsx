import { useNavigate } from "react-router-dom";
import { useTranslation } from "react-i18next";
import { useState } from "react";
import { auth, createUserWithEmailAndPassword } from '../../../firebase';
import { updateProfile } from "firebase/auth";
import { validLettersAndNumbers, validWithNoSpaces } from "../../../utils/InputUtils";
import type { SignUp } from "../../../types/auth";
import GoogleLogin from "../login/GoogleLogin";

const SignupCard = () => {
  const { t } = useTranslation('common');
  const navigate = useNavigate();

  const [formData, setFormData] = useState<SignUp>({ email: "", password: "", repeatPassword: "" });

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

    if (formData.password !== formData.repeatPassword) {
      console.log('password doesnt match')
      return;
    }

    try {
      const cred = await createUserWithEmailAndPassword(auth, formData.email, formData.password);
      await updateProfile(cred.user, { displayName: formData.email.split('@')[0] });
      navigate('/dashboard');
    } catch (err) {
      console.log(err)
    }
  };

  return (
    <>
      <div className="flex justify-center mt-20">
        <fieldset className="fieldset bg-base-200 border-base-300 rounded-box w-xs border p-4">
          <label className="label">{t('signUp.email')}</label>
          <input type="email" className="input" name="email" value={formData.email} onChange={handleEmailChange} placeholder={t('signUp.email')} />

          <label className="label">{t('signUp.password')}</label>
          <input type="password" className="input" name="password" value={formData.password} onChange={handlePasswordChange} placeholder={t('signUp.password')} />

          <label className="label">{t('signUp.repeatPassword')}</label>
          <input type="password" className="input" name="repeatPassword" value={formData.repeatPassword} onChange={handlePasswordChange} placeholder={t('signUp.repeatPassword')} />

          <button className="btn btn-neutral mt-4" onClick={handleSignup}>{t('signUp.signUp')}</button>
        </fieldset>
      </div>
      <GoogleLogin />
    </>
  );
}

export default SignupCard;
