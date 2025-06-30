import { Link, Outlet } from 'react-router-dom';

const PublicLayout = () => {
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
                Home
              </Link>
            </li>
            <li>
              <Link to="/">
                Plans
              </Link>
            </li>
          </ul>
        </div>
        <div className="navbar-end">
          <Link to="/login" className="btn btn-ghost">
            Login
          </Link>
          <Link to="/signup" className='btn btn-ghost ml-1'>
            Sign Up
          </Link>
        </div>
      </div>

      <main>
        <Outlet />
      </main>
    </>
  );
}

export default PublicLayout;
