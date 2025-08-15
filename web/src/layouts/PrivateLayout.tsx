import { Outlet, useNavigate } from 'react-router-dom';
import { IoEarthSharp } from "react-icons/io5";
import { useAuth } from '../context/AuthContext';
import { useTranslation } from 'react-i18next';
import UpgradePlanBanner from '../components/common/UpgradePlanBanner';
import { useProfile } from '../context/ProfileContext';
import { USER_PLAN_ID } from '../constants';

const PrivateLayout = () => {
  const { t, i18n } = useTranslation('common');
  const navigate = useNavigate();
  const { logout } = useAuth();
  const { user } = useProfile();

  const changeLanguage = (lng: string) => {
    i18n.changeLanguage(lng);
  };

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

    
  return (
    <div className="bg-gray-50 min-h-dvh">
      <div className="navbar bg-base-100 shadow-sm">
        <div className="flex-1">
          <a className="btn btn-ghost text-xl">Payment</a>
        </div>
        <div className="flex gap-2">
          <button className='btn btn-outline'>{t('privateLayout.newAlert')}</button>
          <div className="dropdown dropdown-end">
            <div tabIndex={0} role="button" className="btn btn-ghost btn-circle">
              <IoEarthSharp />
            </div>
            <ul
              tabIndex={0}
              className="menu menu-sm dropdown-content bg-base-100 rounded-box z-1 mt-3 w-52 p-2 shadow">
              <li>
                <button onClick={() => changeLanguage('en')}>
                  <span role="img" aria-label="USA flag">🇺🇸 English</span>
                </button>
              </li>
              <li>
                <button onClick={() => changeLanguage('es')}>
                  <span role="img" aria-label="Spain flag">🇪🇸 Español</span>
                </button>
              </li>
            </ul>
          </div>
          <div className="dropdown dropdown-end">
            <div tabIndex={0} role="button" className="btn btn-ghost btn-circle avatar">
              <div className="w-10 rounded-full">
                <img
                  alt="Tailwind CSS Navbar component"
                  src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp" />
              </div>
            </div>
            <ul
              tabIndex={0}
              className="menu menu-sm dropdown-content bg-base-100 rounded-box z-1 mt-3 w-52 p-2 shadow">
              <li>
                <a className="justify-between">
                  {t('privateLayout.profile')}
                </a>
              </li>
              <li><a>{t('privateLayout.settings')}</a></li>
              <li><a onClick={handleLogout}>{t('privateLayout.logout')}</a></li>
            </ul>
          </div>
        </div>
      </div>
      { user?.userPlanId === USER_PLAN_ID.BASIC && <UpgradePlanBanner /> }
      <main>
        <Outlet />
      </main>
    </div>
  );
}

export default PrivateLayout
