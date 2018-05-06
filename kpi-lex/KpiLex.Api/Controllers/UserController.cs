using System;
using KpiLex.Api.Constants;
using KpiLex.Api.Controllers.Abstract;
using KpiLex.Api.Mappers.Abstract;
using KpiLex.Api.Models.Request;
using KpiLex.Api.Models.Response;
using KpiLex.BusinessLogic.Services.Abstract;
using KpiLex.Domain.Models;
using Microsoft.AspNetCore.Mvc;

namespace KpiLex.Api.Controllers
{
    [Route(RouteConstants.UserRoute)]
    public class UserController : BaseRestApiController<UserRequestModel, UserResponseModel, User, Guid>
    {
        public UserController(IService<User, Guid> service, IApiMapper<User, UserResponseModel> responseModelMapper,
            IApiMapper<UserRequestModel, User> requestModelMapper)
            : base(service, responseModelMapper, requestModelMapper)
        {
        }
    }
}