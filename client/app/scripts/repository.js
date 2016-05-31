var modules = []

function getJSON(url) {
    return new Promise(function(resolve, reject) {
        var xhr = new XMLHttpRequest();
        xhr.open('get', url, true);
        xhr.responseType = 'json';
        xhr.onload = function() {
            var status = xhr.status;
            if (status == 200) {
                resolve(xhr.response);
            } else {
                reject(status);
            }
        };
        xhr.send();
    });
};

function getModules() {
    var self = this;
    getJSON('http://localhost:8080/Modules').then(function(data) {
        data.forEach(function(module) {
            module["actions"].forEach(function(element) {
                element["sample"] = JSON.stringify(element["sample"], null, 4);
            }, this);
        }, this);

        modules = data.slice()
        var event = new CustomEvent("modulesReceived", {
            "detail": modules
        });
        window.dispatchEvent(event);
    }, function(status) {
        app.$.toast.text = 'Unable to retrieve modules from Marvin';
        app.$.toast.show();
    });
};

function getModuleByName(name) {
    for (var i in modules) {
        if (modules[i]["name"] == name) {
            return modules[i];
        }
    }
};