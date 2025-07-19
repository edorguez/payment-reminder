import { useNavigate } from "react-router-dom";
import { useAuth } from "../../../context/AuthContext";
import { useTranslation } from "react-i18next";

const LoginCard = () => {
  const { t } = useTranslation('common');
  const navigate = useNavigate();
  const { login } = useAuth();

  const handleLogin = () => {
    login('my-long-token');
    navigate('/dashboard');
  }

  return (
    <>
      <div className="flex justify-center mt-20">
        <fieldset className="fieldset bg-base-200 border-base-300 rounded-box w-xs border p-4">
          <label className="label">{t('login.email')}</label>
          <input type="email" className="input" placeholder={t('login.email')} />

          <label className="label">{t('login.password')}</label>
          <input type="password" className="input" placeholder={t('login.password')} />

          <button className="btn btn-neutral mt-4" onClick={handleLogin}>{t('login.login')}</button>
        </fieldset>
      </div>
    </>
  );
}

export default LoginCard;
