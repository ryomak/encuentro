class Api::V1::Admin::UserPlansController < ApplicationController
  before_action :authenticate_user
  before_action :admin_user?
  before_action :set_user, only: %i[index]

  def index
    plans = @user.plans
    render json: plans
  end

  private
  def set_user
    @user = User.find(params[:user_id])
  end
end
