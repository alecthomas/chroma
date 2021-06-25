wire A , B , C , D , E ; // simple 1 -bit wide wires
wire [8:0] Wide ; // a 9 -bit wide wire
reg I ;

assign A = B & C ; // using a wire with an assign statement

always @ ( B or C ) begin
I = B | C ; // using wires on the right - hand side of an always@
// assignment
end
