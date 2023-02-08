let myEl = document.getElementById("myParragraph")
myEl.classList.add("redText")
document.getElementById("myButton").addEventListener("click", () => {
    // notify("This is my message", "success")
    // notifyModal("title", "<em>Hello world</em>", "success", "My Text for the button")
    let html = `
            <form id="check-availability-form" action="" method="post" novalidate class="need-validation">
                <div class="row">
                    <div class="col-12">
                        <div class="row" id="reservation-dates-modal">
                            <div class="col">
                                <input disabled required class="form-control" type="text" name="start" id="start" placeholder="Arrival">
                            </div>
                            <div class="col">
                                <input disabled required class="form-control" type="text" name="end" id="end" placeholder="Departure">
                            </div>
                        </div>
                    </div>
                </div>
            </form>
            `
    attention.custom({ msg: html, title: "Choose your dates" })
})