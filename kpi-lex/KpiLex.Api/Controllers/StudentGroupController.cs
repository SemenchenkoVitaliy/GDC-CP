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
    [Route(RouteConstants.StudentGroupRoute)]
    public class
        StudentGroupController : BaseRestApiController<StudentRequestModel, StudentGroupResponseModel, StudentGroup, int
        >
    {
        public StudentGroupController(IService<StudentGroup, int> service,
            IApiMapper<StudentGroup, StudentGroupResponseModel> responseModelMapper,
            IApiMapper<StudentRequestModel, StudentGroup> requestModelMapper)
            : base(service, responseModelMapper, requestModelMapper)
        {
        }
    }
}