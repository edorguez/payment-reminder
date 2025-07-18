import { Link, Outlet } from 'react-router-dom';
import { IoEarthSharp } from "react-icons/io5";
import { useTranslation } from 'react-i18next';

const PublicLayout = () => {
  const { t, i18n } = useTranslation('common');

  const changeLanguage = (lng: string) => {
    i18n.changeLanguage(lng);
  };
  return (
    <>
      <div className="navbar bg-base-100 shadow-sm">
        <div className="navbar-start">
          <div className="dropdown">
            <div tabIndex={0} role="button" className="btn btn-ghost lg:hidden">
              <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"> <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M4 6h16M4 12h8m-8 6h16" /> </svg>
            </div>
            <ul
              tabIndex={0}
              className="menu menu-sm dropdown-content bg-base-100 rounded-box z-1 mt-3 w-52 p-2 shadow">
              <li><a>Mobile Item 1</a></li>
              <li><a>Mobile Item 3</a></li>
            </ul>
          </div>
          <Link to="/" className="btn btn-ghost text-xl">
            Payment
          </Link>
        </div>
        <div className="navbar-center hidden lg:flex">
          <ul className="menu menu-horizontal px-1">
            <li>
              <Link to="/">
                {t('home.home')}
              </Link>
            </li>
            <li>
              <Link to="/">
                {t('home.plans')}
              </Link>
            </li>
          </ul>
        </div>
        <div className="navbar-end">
          <Link to="/login" className="btn btn-ghost">
            {t('home.login')}
          </Link>
          <Link to="/signup" className='btn btn-ghost mx-1'>
            {t('home.signUp')}
          </Link>
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
        </div>
      </div>

      <main>
        <Outlet />
      </main>
    </>
  );
}

export default PublicLayout;
