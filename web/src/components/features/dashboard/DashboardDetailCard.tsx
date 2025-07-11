import React from "react";
import type { IconType } from "react-icons";

interface DashboardDetailCardProps {
  icon: IconType;
  title: string;
  total: number;
}

const DashboardDetailCard: React.FC<DashboardDetailCardProps> = ({
  icon: Icon,
  title,
  total
}) => {
  return  (
    <div className="card bg-base-100 sm:w-full  shadow-sm rounded-md">
      <div className="card-body">
        <div className="flex justify-center items-center">
          <Icon className="text-lg text-primary" />
          <h1 className="mx-2">{ title }:</h1>
          <span className="text-xl text-secondary font-bold">{total}</span>
        </div>
      </div>
    </div>
  );
};

export default DashboardDetailCard;
