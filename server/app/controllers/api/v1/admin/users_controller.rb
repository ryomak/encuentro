class Api::V1::Admin::UsersController < ApplicationController
  before_action :authenticate_user
  before_action :admin_user?
  before_action :load_resource
  def index
    users = User.all
    render json: users
  end

  def show
    render json: { user: @user, plans: @user.plans }
  end

  def update
    if @user.update!(user_params)
      render json: @user, status: 200
    else
      render json: { errors: @user.errors }, status:422
    end
  end

  def destroy
    @user.destroy
    head 204
  end

  private

  def load_resource
    case params[:action].to_sym
      when :show
        @user = User.select(:name, :sex, :email, :birthday, :university, :status).find(params[:id]).includes(:plans)
      when :update
        @user = User.select(:name, :sex, :email, :birthday, :university, :status).find(params[:id])
      when :destroy
        @user = User.find(params[:id])
      end
  end

  def user_params
    params.permit(:name, :sex, :email, :password, :password_confirmation, :birthday, :university, :status)
  end
end
