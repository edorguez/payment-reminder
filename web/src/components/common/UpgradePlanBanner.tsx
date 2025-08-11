//import { useTranslation } from "react-i18next";
import { CiCircleInfo } from "react-icons/ci";
import Container from "../ui/Container";
import { IoCloseCircle } from "react-icons/io5";

const UpgradePlanBanner = () => {
  //const { t, i18n } = useTranslation('common');

  return (
    <Container>
      <div className='pt-2'>
        <div role="alert" className="alert alert-vertical sm:alert-horizontal alert-warning alert-soft">
          <div className="text-warning text-2xl">
            <CiCircleInfo />
          </div>
          <div>
            <h3 className="font-bold">New message!</h3>
            <div className="text-xs">You have 1 unread message</div>
          </div>
          <button className="btn text-lg text-white bg-secondary">
            <IoCloseCircle />
          </button>
        </div>
      </div>
    </Container>
  );
}

export default UpgradePlanBanner;
