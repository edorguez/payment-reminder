import { useNavigate } from "react-router-dom";
import { useAuth } from "../../../context/AuthContext";
import { useTranslation } from "react-i18next";

const SignupCard = () => {
  const { t } = useTranslation('common');
  const navigate = useNavigate();
  const { login } = useAuth();

  const handleSignup = () => {
    login('my-long-token');
    navigate('/dashboard');

  }
  return (
    <>
      <div className="flex justify-center mt-20">
        <fieldset className="fieldset bg-base-200 border-base-300 rounded-box w-xs border p-4">
          <label className="label">{t('signUp.email')}</label>
          <input type="email" className="input" placeholder={t('signUp.email')} />

          <label className="label">{t('signUp.password')}</label>
          <input type="password" className="input" placeholder={t('signUp.password')} />

          <label className="label">{t('signUp.repeatPassword')}</label>
          <input type="password" className="input" placeholder={t('signUp.repeatPassword')} />

          <button className="btn btn-neutral mt-4" onClick={handleSignup}>{t('signUp.signUp')}</button>
        </fieldset>
      </div>
    </>
  );
}

export default SignupCard;
