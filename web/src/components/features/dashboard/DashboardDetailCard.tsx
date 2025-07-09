const DashboardDetailCard = () => {
  return  (
    <div className="card bg-base-100 sm:w-full  shadow-sm rounded-md">
      <div className="card-body">
        <h2 className="card-title">Card title!</h2>
        <p>A card component has a figure, a body part, and inside body there are title and actions parts</p>
        <div className="card-actions justify-end">
          <button className="btn btn-primary">Buy Now</button>
        </div>
      </div>
    </div>
  );
};

export default DashboardDetailCard;
