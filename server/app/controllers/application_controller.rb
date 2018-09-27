class ApplicationController < ActionController::API
  include Knock::Authenticable

  private
  def admin_user?
    if current_user.user?
      render json: { message:  "管理者権限がありません"}, status: 500
    end
  end
end
