import React from "react";

interface DashboardDetailCardProps {
  iconPath: string;
  title: string;
  total: number;
}

const DashboardDetailCard: React.FC<DashboardDetailCardProps> = ({
  iconPath,
  title,
  total
}) => {
  return  (
    <div className="card bg-base-100 sm:w-full  shadow-sm rounded-md">
      <div className="card-body">
        <div className="flex justify-center items-center">
          <img className="w-4" src={iconPath} />
          <h1 className="mx-2">{ title }:</h1>
          <span className="text-xl text-secondary font-bold">{total}</span>
        </div>
      </div>
    </div>
  );
};

export default DashboardDetailCard;
