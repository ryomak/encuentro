class Api::V1::Admin::PlansController < ApplicationController
  before_action :authenticate_user
  before_action :admin_user?
  before_action :set_plan, only: %i[show update destroy]

  def index
    plans = Plan.all
    render json: plans
  end

  def show
    render json: @plan
  end

  def update
    if @plan.update(plan_params)
      render json: @plan, status: 200
    else
      render json: { errors: plan.errors }, status:422
    end
  end

  def destroy
    @plan.destroy
    head 204
  end

  private

  def set_plan
    @plan = Plan.find(params[:id])
  end

  def plan_params
    params.permit(:date, :place, :drink, :course, :status)
  end

end
