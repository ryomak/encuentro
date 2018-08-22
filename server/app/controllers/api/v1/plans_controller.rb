class Api::V1::PlansController < ApplicationController
  before_action :authoritative
  before_action :set_plan, only: %i[show update destroy]

  def index
    @plans = set_user.plans
    render json: @plans
  end

  def show
    render json: @plan
  end

  def create
    plan = Plan.new(plan_params)
    plan.user = set_user
    if plan.save
      render json: plan, status: 201
    else
      render json: { errors: plan.errors }, status: 422
    end
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

  def set_user
    User.find(current_user.id)
  end

  def set_plan
    @plan = Plan.find(params[:id])
  end

  def plan_params
    params.permit(:date, :place, :drink, :course)
  end
end
