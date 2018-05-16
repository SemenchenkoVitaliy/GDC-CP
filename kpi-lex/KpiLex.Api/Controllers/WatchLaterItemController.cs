using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using KpiLex.Api.Constants;
using KpiLex.Api.Mappers.Abstract;
using KpiLex.Api.Models.Request;
using KpiLex.Api.Models.Response;
using KpiLex.BusinessLogic.Services.Abstract;
using KpiLex.Domain.Models;
using Microsoft.AspNetCore.Mvc;

namespace KpiLex.Api.Controllers
{
    [Route(RouteConstants.WatchLaterRoute)]
    public class WatchLaterItemController : Controller
    {
        private readonly IWatchLaterService _watchLaterItemService;
        private readonly IApiMapper<WatchLaterItem, WatchLaterItemResponseModel> _responseModelMapper;
        private readonly IApiMapper<WatchLaterItemRequestModel, WatchLaterItem> _requestModelMapper;

        public WatchLaterItemController(
            IWatchLaterService watchLaterItemService,
            IApiMapper<WatchLaterItem, WatchLaterItemResponseModel> responseModelMapper,
            IApiMapper<WatchLaterItemRequestModel, WatchLaterItem> requestModelMapper)
        {
            _watchLaterItemService = watchLaterItemService;
            _responseModelMapper = responseModelMapper;
            _requestModelMapper = requestModelMapper;
        }
        
        [HttpGet]
        public async Task<IReadOnlyCollection<WatchLaterItemResponseModel>> Get([FromRoute] int studentId)
        {
            var courseComments = await _watchLaterItemService.GetWatchLaterItems(studentId);
            var responseModels = courseComments.Select(c => _responseModelMapper.Map(c)).ToList();
            return responseModels;
        }
        
        [HttpPost]
        public async Task<WatchLaterItemResponseModel> Add(
            [FromRoute] int courseId,
            [FromBody] WatchLaterItemRequestModel courseCommentRequestModel)
        {
            var model = _requestModelMapper.Map(courseCommentRequestModel);
            model = await _watchLaterItemService.AddAsync(model);
            return _responseModelMapper.Map(model);
        }
        
        [HttpPut]
        public async Task<IActionResult> Update(
            [FromRoute] int courseId,
            [FromBody] WatchLaterItemRequestModel courseCommentRequestModel)
        {
            var model = _requestModelMapper.Map(courseCommentRequestModel);
            await _watchLaterItemService.UpdateAsync(model);
            return Ok();
        }
        
        [HttpDelete, Route("{watchLaterItemId}")]
        public async Task<IActionResult> Delete(
            [FromRoute] int courseId,
            [FromRoute] Guid watchLaterItemId)
        {
            await _watchLaterItemService.RemoveAsync(watchLaterItemId);
            return Ok();
        }
    }
}