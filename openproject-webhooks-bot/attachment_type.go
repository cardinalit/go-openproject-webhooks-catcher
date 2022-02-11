package openproject_webhooks_bot

import (
   "time"
)

type Attachment struct {
   Action     string `json:"action"`
   Attachment struct {
      Type        string `json:"_type"`
      ID          int    `json:"id"`
      FileName    string `json:"fileName"`
      FileSize    int    `json:"fileSize"`
      Description struct {
         Format string      `json:"format"`
         Raw    interface{} `json:"raw"`
         HTML   string      `json:"html"`
      } `json:"description"`
      ContentType string `json:"contentType"`
      Digest      struct {
         Algorithm string `json:"algorithm"`
         Hash      string `json:"hash"`
      } `json:"digest"`
      CreatedAt time.Time `json:"createdAt"`
      Embedded  struct {
         Author struct {
            Type        string      `json:"_type"`
            ID          int         `json:"id"`
            Name        string      `json:"name"`
            CreatedAt   time.Time   `json:"createdAt"`
            UpdatedAt   time.Time   `json:"updatedAt"`
            Login       string      `json:"login"`
            Admin       bool        `json:"admin"`
            FirstName   string      `json:"firstName"`
            LastName    string      `json:"lastName"`
            Email       string      `json:"email"`
            Avatar      string      `json:"avatar"`
            Status      string      `json:"status"`
            IdentityURL interface{} `json:"identityUrl"`
            Language    string      `json:"language"`
            Links       struct {
               Self struct {
                  Href  string `json:"href"`
                  Title string `json:"title"`
               } `json:"self"`
               Memberships struct {
                  Href  string `json:"href"`
                  Title string `json:"title"`
               } `json:"memberships"`
               ShowUser struct {
                  Href string `json:"href"`
                  Type string `json:"type"`
               } `json:"showUser"`
               UpdateImmediately struct {
                  Href   string `json:"href"`
                  Title  string `json:"title"`
                  Method string `json:"method"`
               } `json:"updateImmediately"`
               Lock struct {
                  Href   string `json:"href"`
                  Title  string `json:"title"`
                  Method string `json:"method"`
               } `json:"lock"`
               Delete struct {
                  Href   string `json:"href"`
                  Title  string `json:"title"`
                  Method string `json:"method"`
               } `json:"delete"`
               AuthSource struct {
                  Href  string `json:"href"`
                  Title string `json:"title"`
               } `json:"auth_source"`
            } `json:"_links"`
         } `json:"author"`
         Container struct {
            Type        string `json:"_type"`
            ID          int    `json:"id"`
            LockVersion int    `json:"lockVersion"`
            Subject     string `json:"subject"`
            Description struct {
               Format string `json:"format"`
               Raw    string `json:"raw"`
               HTML   string `json:"html"`
            } `json:"description"`
            ScheduleManually     bool        `json:"scheduleManually"`
            StartDate            interface{} `json:"startDate"`
            DueDate              interface{} `json:"dueDate"`
            DerivedStartDate     interface{} `json:"derivedStartDate"`
            DerivedDueDate       interface{} `json:"derivedDueDate"`
            EstimatedTime        interface{} `json:"estimatedTime"`
            DerivedEstimatedTime interface{} `json:"derivedEstimatedTime"`
            PercentageDone       int         `json:"percentageDone"`
            CreatedAt            time.Time   `json:"createdAt"`
            UpdatedAt            time.Time   `json:"updatedAt"`
            Links                struct {
               Attachments struct {
                  Href string `json:"href"`
               } `json:"attachments"`
               AddAttachment struct {
                  Href   string `json:"href"`
                  Method string `json:"method"`
               } `json:"addAttachment"`
               Self struct {
                  Href  string `json:"href"`
                  Title string `json:"title"`
               } `json:"self"`
               Update struct {
                  Href   string `json:"href"`
                  Method string `json:"method"`
               } `json:"update"`
               Schema struct {
                  Href string `json:"href"`
               } `json:"schema"`
               UpdateImmediately struct {
                  Href   string `json:"href"`
                  Method string `json:"method"`
               } `json:"updateImmediately"`
               Delete struct {
                  Href   string `json:"href"`
                  Method string `json:"method"`
               } `json:"delete"`
               LogTime struct {
                  Href  string `json:"href"`
                  Title string `json:"title"`
               } `json:"logTime"`
               Move struct {
                  Href  string `json:"href"`
                  Type  string `json:"type"`
                  Title string `json:"title"`
               } `json:"move"`
               Copy struct {
                  Href  string `json:"href"`
                  Title string `json:"title"`
               } `json:"copy"`
               Pdf struct {
                  Href  string `json:"href"`
                  Type  string `json:"type"`
                  Title string `json:"title"`
               } `json:"pdf"`
               Atom struct {
                  Href  string `json:"href"`
                  Type  string `json:"type"`
                  Title string `json:"title"`
               } `json:"atom"`
               AvailableRelationCandidates struct {
                  Href  string `json:"href"`
                  Title string `json:"title"`
               } `json:"availableRelationCandidates"`
               CustomFields struct {
                  Href  string `json:"href"`
                  Type  string `json:"type"`
                  Title string `json:"title"`
               } `json:"customFields"`
               ConfigureForm struct {
                  Href  string `json:"href"`
                  Type  string `json:"type"`
                  Title string `json:"title"`
               } `json:"configureForm"`
               Activities struct {
                  Href string `json:"href"`
               } `json:"activities"`
               AvailableWatchers struct {
                  Href string `json:"href"`
               } `json:"availableWatchers"`
               Relations struct {
                  Href string `json:"href"`
               } `json:"relations"`
               Revisions struct {
                  Href string `json:"href"`
               } `json:"revisions"`
               Watchers struct {
                  Href string `json:"href"`
               } `json:"watchers"`
               AddWatcher struct {
                  Href    string `json:"href"`
                  Method  string `json:"method"`
                  Payload struct {
                     User struct {
                        Href string `json:"href"`
                     } `json:"user"`
                  } `json:"payload"`
                  Templated bool `json:"templated"`
               } `json:"addWatcher"`
               RemoveWatcher struct {
                  Href      string `json:"href"`
                  Method    string `json:"method"`
                  Templated bool   `json:"templated"`
               } `json:"removeWatcher"`
               AddRelation struct {
                  Href   string `json:"href"`
                  Method string `json:"method"`
                  Title  string `json:"title"`
               } `json:"addRelation"`
               AddChild struct {
                  Href   string `json:"href"`
                  Method string `json:"method"`
                  Title  string `json:"title"`
               } `json:"addChild"`
               ChangeParent struct {
                  Href   string `json:"href"`
                  Method string `json:"method"`
                  Title  string `json:"title"`
               } `json:"changeParent"`
               AddComment struct {
                  Href   string `json:"href"`
                  Method string `json:"method"`
                  Title  string `json:"title"`
               } `json:"addComment"`
               PreviewMarkup struct {
                  Href   string `json:"href"`
                  Method string `json:"method"`
               } `json:"previewMarkup"`
               TimeEntries struct {
                  Href  string `json:"href"`
                  Title string `json:"title"`
               } `json:"timeEntries"`
               Ancestors []interface{} `json:"ancestors"`
               Category  struct {
                  Href interface{} `json:"href"`
               } `json:"category"`
               Type struct {
                  Href  string `json:"href"`
                  Title string `json:"title"`
               } `json:"type"`
               Priority struct {
                  Href  string `json:"href"`
                  Title string `json:"title"`
               } `json:"priority"`
               Project struct {
                  Href  string `json:"href"`
                  Title string `json:"title"`
               } `json:"project"`
               Status struct {
                  Href  string `json:"href"`
                  Title string `json:"title"`
               } `json:"status"`
               Author struct {
                  Href  string `json:"href"`
                  Title string `json:"title"`
               } `json:"author"`
               Responsible struct {
                  Href interface{} `json:"href"`
               } `json:"responsible"`
               Assignee struct {
                  Href interface{} `json:"href"`
               } `json:"assignee"`
               Version struct {
                  Href interface{} `json:"href"`
               } `json:"version"`
               Parent struct {
                  Href  interface{} `json:"href"`
                  Title interface{} `json:"title"`
               } `json:"parent"`
               CustomActions []struct {
                  Href  string `json:"href"`
                  Title string `json:"title"`
               } `json:"customActions"`
               Github struct {
                  Href  string `json:"href"`
                  Title string `json:"title"`
               } `json:"github"`
               GithubPullRequests struct {
                  Href  string `json:"href"`
                  Title string `json:"title"`
               } `json:"github_pull_requests"`
               ConvertBCF struct {
                  Href    string `json:"href"`
                  Title   string `json:"title"`
                  Payload struct {
                     ReferenceLinks []string `json:"reference_links"`
                  } `json:"payload"`
                  Method string `json:"method"`
               } `json:"convertBCF"`
               CustomField5 struct {
                  Title string `json:"title"`
                  Href  string `json:"href"`
               } `json:"customField5"`
            } `json:"_links"`
         } `json:"container"`
      } `json:"_embedded"`
      Links struct {
         Self struct {
            Href  string `json:"href"`
            Title string `json:"title"`
         } `json:"self"`
         Author struct {
            Href  string `json:"href"`
            Title string `json:"title"`
         } `json:"author"`
         Container struct {
            Href  string `json:"href"`
            Title string `json:"title"`
         } `json:"container"`
         StaticDownloadLocation struct {
            Href string `json:"href"`
         } `json:"staticDownloadLocation"`
         DownloadLocation struct {
            Href string `json:"href"`
         } `json:"downloadLocation"`
         Delete struct {
            Href   string `json:"href"`
            Method string `json:"method"`
         } `json:"delete"`
      } `json:"_links"`
   } `json:"attachment"`
}
