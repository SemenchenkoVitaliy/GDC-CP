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
    [Route(RouteConstants.TeacherRoute)]
    public class TeacherController : BaseRestApiController<TeacherRequestModel, TeacherResponseModel, Teacher, long>
    {
        public TeacherController(IService<Teacher, long> service,
            IApiMapper<Teacher, TeacherResponseModel> responseModelMapper,
            IApiMapper<TeacherRequestModel, Teacher> requestModelMapper) : base(service, responseModelMapper, requestModelMapper)
        {
        }
    }
}