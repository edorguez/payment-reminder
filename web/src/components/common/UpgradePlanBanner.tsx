import { useTranslation } from "react-i18next";
import { CiCircleInfo } from "react-icons/ci";
import Container from "../ui/Container";
import { IoCloseCircle } from "react-icons/io5";
import { useState } from "react";

const UpgradePlanBanner = () => {
  const { t } = useTranslation('common');
  const [showMessage, setShowMessage] = useState<boolean>(true);

  if(!showMessage)
    return <></>;

  return (
    <>
      <Container>
        <div className='pt-2'>
          <div role="alert" className="alert alert-vertical sm:alert-horizontal alert-warning alert-soft">
            <div className="text-warning text-2xl">
              <CiCircleInfo />
            </div>
            <div>
              <h3 className="font-bold">{t('dashboard.upgradePlanTitle')}</h3>
              <div className="text-xs">{t('dashboard.upgradePlanDescription')}</div>
            </div>
            <button className="btn text-lg text-white bg-secondary" onClick={() => { setShowMessage(false) }} >
              <IoCloseCircle />
            </button>
          </div>
        </div>
      </Container>
    </>
  );
}

export default UpgradePlanBanner;
