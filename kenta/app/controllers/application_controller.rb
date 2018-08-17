class ApplicationController < ActionController::API
  include Knock::Authenticable

  private
  def admin_user?
    if !current_user.status
      render json: { text:  "reject"}
    end
  end
end
