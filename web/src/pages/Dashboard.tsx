import { useTranslation } from "react-i18next";
import DashboardDetailCard from "../components/features/dashboard/DashboardDetailCard";
import DashboardTable from "../components/features/dashboard/DashboardTable";
import Container from "../components/ui/Container";

const Dashboard = () => {
  const { t } = useTranslation('common');

  return (
    <Container>
      <div className="py-4 grid sm:grid-cols-3 gap-3">
        <DashboardDetailCard iconPath="/images/dashboard/bell.png" title={t('dashboard.upcomingAlerts')} total={15} />
        <DashboardDetailCard iconPath="/images/dashboard/bell-sleep.png" title={t('dashboard.pausedAlerts')} total={5} />
        <DashboardDetailCard iconPath="/images/dashboard/list.png" title={t('dashboard.totalAlerts')} total={100} />
      </div>
      <DashboardTable />
    </Container>
  )
}

export default Dashboard;
