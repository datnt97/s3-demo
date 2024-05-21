package response

type UploadResp struct {
	Url string `json:"url"`
}

// func AttachmentMapToResponses(items *attachment.UploadAttachmentResponse) []*AttachmentResp {
// 	if items.GetAttachments() == nil {
// 		return nil
// 	}

// 	var results []*AttachmentResp
// 	for _, v := range items.GetAttachments() {
// 		results = append(results, &AttachmentResp{
// 			Url: v.GetUrl(),
// 		})
// 	}
// 	return results
// }
