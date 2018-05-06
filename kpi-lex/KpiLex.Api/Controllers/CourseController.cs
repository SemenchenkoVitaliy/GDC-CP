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
    [Route(RouteConstants.CourseRoute)]
    public class CourseController : BaseRestApiController<
        CourseCommentRequestModel, 
        CourseCommentResponseModel, 
        CourseComment, 
        long>
    {
        public CourseController(
            IService<CourseComment, long> service, 
            IApiMapper<CourseComment, CourseCommentResponseModel> responseModelMapper, 
            IApiMapper<CourseCommentRequestModel, CourseComment> requestModelMapper) 
            : base(service, responseModelMapper, requestModelMapper)
        {
        }
    }
}